
drop DATABASE if EXISTS bd_aveonline;
CREATE DATABASE bd_aveonline;
COMMENT on DATABASE bd_aveonline is 'base de datos para segmento de medicamentos';

drop table if exists medicamento;
create table medicamento(
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY key, 
  nombre varchar(80)  not null,
  precio float NOT NULL CHECK(precio>0),
  ubicacion varchar(50) not NULL
);

drop table if exists promocion;
create table promocion(
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY key,
  descripcion varchar(100) not null,
  porcentaje float NOT NULL CHECK(porcentaje>0 and porcentaje<70),
  fecha_inicio date not NULL,
  fecha_fin date not  NULL,
  CHECK(fecha_inicio<fecha_fin)
);

-- trigger para tener las fechas de promociones no choquen
CREATE OR REPLACE FUNCTION proteger_datos () RETURNS trigger AS
$$
BEGIN
   if exists(select*from promocion 
   where new.fecha_inicio BETWEEN fecha_inicio and fecha_fin or 
   new.fecha_fin BETWEEN fecha_inicio and fecha_fin) then
     raise exception 'se esta registrando un choque de promociones';
return null;
   end if;
   return new;
END;
$$ LANGUAGE 'plpgsql';


CREATE  TRIGGER promociones_choques BEFORE Insert 
    ON promocion FOR EACH ROW 
    EXECUTE PROCEDURE proteger_datos();

drop table if exists factura;
create table factura(
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY key,
  fecha_crear date DEFAULT current_timestamp,
  pago_total float NOT NULL check(pago_total>0),
  promocion_id integer NOT NULL REFERENCES promocion
);

drop table if exists factura_items;
create table factura_items(
  factura_id integer references factura,
  medicamento_id integer references medicamento 
);
