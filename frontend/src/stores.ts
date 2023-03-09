import { writable } from "svelte/store";
import type { main } from "wailsjs/go/models";

export const sourceFolder = writable<string | null>(null);
export const targetFolder = writable<string | null>(null);
export const isReady = writable<boolean>(false);

export const sourceFiles = writable<main.Image[]>([]);

export const uiColumns = writable(5);
export const uiHideCopied = writable(false);

export const uiChunkPage = writable(0);
export const uiChunkSize = writable(200);

export const loading = writable<boolean>(false);

export const logrocketRecording = writable<boolean>(false);

export const showPrivacyPolicy = writable<boolean>(false);
