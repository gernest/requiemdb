package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()

	a, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	a = bytes.ReplaceAll(a, []byte("export declare"), []byte("export "))
	a = bytes.ReplaceAll(a, []byte("declare class"), []byte("export class"))
	a = bytes.ReplaceAll(a, []byte("declare interface"), []byte("export interface"))
	a = bytes.ReplaceAll(a, []byte("declare type"), []byte("export type"))
	a = bytes.ReplaceAll(a, []byte("declare const"), []byte("export const"))
	a = bytes.ReplaceAll(a, []byte("declare enum"), []byte("export enum"))
	a = bytes.ReplaceAll(a, []byte("declare abstract"), []byte("export abstract"))
	mod := fmt.Sprintf("declare module  '@requiemdb/rq'{\n %s\n}", string(a))
	a, _ = json.Marshal(map[string]string{
		"requiem": mod,
	})
	// os.WriteFile("mod.d.ts", []byte(mod), 0600)
	err = os.WriteFile(flag.Arg(1), a, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
