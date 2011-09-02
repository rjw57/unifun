#!/bin/bash
gomake -C pkg/uctricks "$@"
gomake -C cmd "$@"
