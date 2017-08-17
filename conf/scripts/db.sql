create database btchist;

CREATE USER 'btchist'@'localhost'
IDENTIFIED BY '';


grant insert, select, update on btchist.* to 'btchist'@'localhost' ;

create table btchist.site (
id_site integer not null AUTO_INCREMENT 
,nm_site varchar(50) not null
,link_site varchar(150) not null
,id_mode integer not null 
,PRIMARY KEY (id_site)
);

ALTER TABLE btchist.site 
ADD CONSTRAINT fk_site_mode
FOREIGN KEY ( id_mode ) 
REFERENCES btchist.mode(id_mode);


create table btchist.mode (
id_mode integer not null auto_increment
,nm_mode varchar(25)
,primary key (id_mode)
);

insert into btchist.mode (nm_mode) values ("scrap");
insert into btchist.mode (nm_mode) values ("json");

insert into btchist.site (nm_site, link_site, id_mode) 
    values ("Dolar Hoje", "http://dolarhoje.com/bitcoin-hoje/", 1);
insert into btchist.site (nm_site, link_site, id_mode) 
    values ("Blockchain", "https://blockchain.info/pt/ticker", 2);

create table btchist.currency (
id_currency integer not null auto_increment
,nm_currency varchar(30)
,symbol varchar(3)
,primary key (id_currency) 
);

create table btchist.var_btc_hist (
id_var_btc_hist integer not null auto_increment
,var_btc decimal(15,2)
,id_currency integer not null
,id_site integer not null 
,primary key (id_var_btc_hist)
,constraint fk_var_btc_hist_site foreign key (id_site) references btchist.site (id_site)
,constraint fk_var_btc_hist_currency foreign key (id_currency) references btchist.currency (id_currency)
);

insert into btchist.currency (nm_currency, symbol) values ("Real", "R$");
insert into btchist.currency (nm_currency, symbol) values ("Dolar", "$");
