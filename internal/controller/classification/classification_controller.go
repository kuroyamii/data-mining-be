package classificationController

import (
	"bytes"
	"datamining-be/internal/model"
	"datamining-be/pkg/config"
	baseResponse "datamining-be/pkg/response"
	"datamining-be/pkg/utilities"
	"fmt"
	"log"
	"net/http"
	"strings"

	classificationService "datamining-be/internal/service/classification"

	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type classificationController struct {
	router *mux.Router
	cfg    config.Config
	cs     classificationService.ClassificationService
}

func (cc classificationController) handleClassification(rw http.ResponseWriter, r *http.Request) {

	file, _, err := sicgolib.ParseImageFile(r, "file", 100<<20)
	ik, err := sicgolib.CreateNewImagekitClient(cc.cfg.PublicKey, cc.cfg.PrivateKey)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
	}
	res, err := sicgolib.UploadToImagekit(r.Context(), ik, cc.cfg.PublicKey, cc.cfg.PrivateKey, file, "file", "/datamining")
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
	}

	// to flask API
	bodyText := fmt.Sprintf(`{
		"image_url": "%v"
		}`, res.URL)
	jsonBody := []byte(bodyText)
	resp, err := http.Post("http://localhost:8181/classify", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
	}

	respon := model.ClassifyResponseFromFlask{}

	err = respon.FromJSON(resp.Body)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return
	}

	responseData := respon.Data.ResNet
	responseDataVGG := respon.Data.VGG16
	temp := strings.Split(responseData, "___")
	temp2 := strings.Split(responseDataVGG, "___")
	name := temp[0]
	nameVGG := temp2[0]

	disease := temp[1]
	diseaseVGG := temp2[1]
	condition := ""
	conditionVGG := ""
	if disease == "healthy" {
		condition = "Healthy"
	} else {
		condition = "Unhealthy"
	}
	if diseaseVGG == "healthy" {
		conditionVGG = "Healthy"
	} else {
		conditionVGG = "Unhealthy"
	}

	err = cc.cs.InsertImagePath(r.Context(), res.URL, name, condition, disease)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return
	}

	resMapResNet := map[string]string{
		"name":      strings.ReplaceAll(name, "_", " "),
		"condition": strings.ReplaceAll(condition, "_", " "),
		"disease":   strings.ReplaceAll(disease, "_", " "),
	}
	nameVGG = strings.ReplaceAll(nameVGG, "_", " ")
	conditionVGG = strings.ReplaceAll(conditionVGG, "_", " ")
	diseaseVGG = strings.ReplaceAll(diseaseVGG, "_", " ")

	resMapVGG := map[string]string{
		"name":      strings.ReplaceAll(nameVGG, "_", " "),
		"condition": strings.ReplaceAll(conditionVGG, "_", " "),
		"disease":   strings.ReplaceAll(diseaseVGG, "_", " "),
	}
	resMap := map[string]map[string]string{
		"resnet": resMapResNet,
		"vgg16":  resMapVGG,
	}

	baseResponse.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, resMap).ToJSON(rw)
}

func (cc classificationController) InitializeEndpoints() {
	cc.router.HandleFunc("/classify", cc.handleClassification).Methods("POST")
}

func NewClassificationController(router *mux.Router, cfg config.Config, cs classificationService.ClassificationService) classificationController {
	return classificationController{
		router: router,
		cfg:    cfg,
		cs:     cs,
	}
}
