import { writable } from "svelte/store";
import type { main } from "wailsjs/go/models";

export const sourceFolder = writable(null);
export const targetFolder = writable(null);

export const sourceFiles = writable<main.Image[]>([]);

export type FileBase64Index = {
  [key: string]: string;
};

// Path as key, base64 data as value
export const fileBase64Index = writable<FileBase64Index>({});

export const uiColumns = writable(5);
export const uiHideCopied = writable(false);

export const uiChunkPage = writable(0);
export const uiChunkSize = writable(200);
