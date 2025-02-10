#!/bin/bash
docker build -t blourgdb .
docker run -p 5432:5432 -v ~/blourg/data:/var/lib/postgresql/data blourgdb