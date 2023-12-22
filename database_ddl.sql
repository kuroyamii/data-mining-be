use datamining;

create table diseases(
	id int not null auto_increment,
	name text,
	primary key(id)
);


create table plants(
	id int not null auto_increment,
	name text,
	primary key(id)
);


create table conditions(
	id int not null auto_increment,
	name text,
	primary key(id)
);


create table plant_disease_details(
	id int not null auto_increment,
	plant_id int not null,
	disease_id int not null,
	primary key(id),
	foreign key (plant_id) references plants(id) on delete cascade on update cascade,
	foreign key (disease_id) references diseases(id) on delete cascade on update cascade
);


create table plant_images(
	id int not null auto_increment,
	file_name text,
	plant_disease_detail_id int not null,
	condition_id int not null,
	primary key(id),
	foreign key (plant_disease_detail_id) references plant_disease_details(id) on delete cascade on update cascade,
	foreign key (condition_id) references conditions(id) on delete cascade on update cascade
);



insert into conditions(name) values("healthy"),('unhealthy');
insert into plants(name) values
("Apple"),
("Cherry_(including_sour)"),
("Corn_(maize)"),
("Grape"),
("Peach"),
("Pepper,_bell"),
("Potato"),
("Strawberry"),
("Tomato");


insert into diseases(name) values
('healthy'),
('Apple_scab'),
('Black_rot'),
('Cedar_apple_rust'),
('Powdery_mildew'),
('Cercospora_leaf_spot Gray_leaf_spot'),
('Common_rust_'),
('Northern_Leaf_Blight'),
('Esca_(Black_Measles)'),
('Leaf_blight_(Isariopsis_Leaf_Spot)'),
('Bacterial_spot'),
('Early_blight'),
('Late_blight'),
('Leaf_scorch'),
('Leaf_Mold'),
('Septoria_leaf_spot'),
('Spider_mites Two-spotted_spider_mite'),
('Target_Spot'),
('Tomato_mosaic_virus'),
('Tomato_Yellow_Leaf_Curl_Virus');

-- select * from plants;
-- select * from conditions;
-- select * from diseases;

-- select p.name,d.name from plant_disease_details pdd 
-- inner join plants p on p.id = pdd.plant_id
-- inner join diseases d on d.id = pdd.disease_id ;


insert into plant_disease_details(plant_id, disease_id)
values
(1,2),
(1,3),
(1,4),
(1,1),
(2,1),
(2,5),
(3,6),
(3,7),
(3,1),
(3,8),
(4,3),
(4,9),
(4,1),
(4,10),
(5,11),
(5,1),
(6,11),
(6,1),
(7,12),
(7,1),
(7,13),
(8,1),
(8,14),
(9,11),
(9,12),
(9,1),
(9,13),
(9,15),
(9,16),
(9,17),
(9,18),
(9,19),
(9,20);


