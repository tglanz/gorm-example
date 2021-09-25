drop database if exists playground;
create database playground;
use playground;

create table specs (
	spec_id integer,
    primary key (spec_id)
);

create table commands (
	command_name integer,
    spec_id integer,
    primary key (command_name, spec_id),
    foreign key fk_commands_specs (spec_id) references specs(spec_id)
);

create table execution_contexts (
	run_id integer,
    spec_id integer,
    primary key (run_id, spec_id),
    foreign key fk_execution_contexts (spec_id) references specs(spec_id)
);

create table command_results (
	command_name integer,
    spec_id integer,
    run_id integer,
    primary key (command_name, spec_id, run_id),
    foreign key fk_command_results_commands (command_name, spec_id) references commands(command_name, spec_id),
    foreign key fk_command_results_execution_contexts (spec_id, run_id) references execution_contexts(spec_id, run_id)
);
