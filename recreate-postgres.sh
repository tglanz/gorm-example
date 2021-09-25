#!/bin/bash

sudo -iu postgres psql -c "drop database local" && sudo -iu postgres psql -c "create database local"
