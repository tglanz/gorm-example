drop table if exists command_results;
drop table if exists execution_contexts;
drop table if exists commands;
drop table if exists specs;

create table specs (
	spec_id integer,
    primary key (spec_id)
);

create table commands (
	command_name integer,
    spec_id integer,
    primary key (command_name, spec_id),
    constraint fk_commands_specs foreign key (spec_id) references specs(spec_id)
);

create table execution_contexts (
	run_id integer,
    spec_id integer,
    primary key (run_id, spec_id),
    constraint fk_execution_contexts foreign key (spec_id) references specs(spec_id)
);

create table command_results (
	command_name integer,
    spec_id integer,
    run_id integer,
    primary key (command_name, spec_id, run_id),
    constraint fk_command_results_commands foreign key (command_name, spec_id) references commands(command_name, spec_id),
    constraint fk_command_results_execution_contexts foreign key (spec_id, run_id) references execution_contexts(spec_id, run_id)
);
