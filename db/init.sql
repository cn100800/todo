create table IF NOT EXISTS todo (
   id integer  PRIMARY KEY ,
   title varchar(50),
   content text,
   step text,
   url text,
   branch varchar(50),
   start_time integer,
   end_time integer,
   create_time integer,
   update_time integer
);
