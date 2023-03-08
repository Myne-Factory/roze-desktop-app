import { writable } from "svelte/store";
import type { main } from "wailsjs/go/models";

export const sourceFolder = writable<string | null>();
export const targetFolder = writable<string | null>();

export const sourceFiles = writable<main.Image[]>([]);

export const uiColumns = writable(5);
export const uiHideCopied = writable(false);

export const uiChunkPage = writable(0);
export const uiChunkSize = writable(500);

export const loading = writable<boolean>(false);
