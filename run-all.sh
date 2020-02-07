#!/usr/bin/env bash
cp crypto-tool.yml ~/.teamocil/
sed -i'*.yml' s%CURR_DIR%$(pwd)%g ~/.teamocil/crypto-tool.yml
itermocil crypto-tool