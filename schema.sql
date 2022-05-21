drop table if exists students;
drop table if exists departments;
drop table if exists staffs;
drop table if exists classes;

create table students (
   id serial not null primary key,
   firstname varchar not null,
   lastname varchar not null,
   email varchar not null,
   admission_number varchar not null,
   year int not null
);

create table classes (
   id serial not null primary key,
   title varchar not null,
   code varchar not null,
   unit int not null,
   semester int not null,
   location varchar not null
);

create table departments (
   id serial not null primary key,
   name varchar not null,
   code varchar not null,
   telephone varchar not null,

   foreign key (user_id) references users (id)
);

create table staffs (
   id serial not null primary key,
   firstname varchar not null,
   lastname varchar not null,
   email varchar not null,
   telephone varchar not null,
   salary bigint not null,
);

create table classes_students (
   class_id int not null,
   student_id int not null,

   primary key (class_id, student_id),
   foreign key (student_id) references students (id),
   foreign key (class_id) references classes (id)
);

create table classes_instructors (
   class_id int not null,
   staff_id int not null,

   primary key (class_id, staff_id),
   foreign key (staff_id) references staffs (id),
   foreign key (class_id) references classes (id)
);