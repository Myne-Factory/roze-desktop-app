import App from "./App.svelte";
import "./style.scss";
import { GetPage, SavePage } from "../wailsjs/go/main/App";
import { sourceFiles, uiChunkPage } from "./stores";
import { get } from "svelte/store";
const app = new App({
  target: document.getElementById("app"),
});

sourceFiles.subscribe(async (files) => {
  if (files.length > 0) {
    const page = await GetPage(files[0].fullPath);
    if (page) {
      uiChunkPage.set(page);
    }
  }
});

uiChunkPage.subscribe(async (number) => {
  if (number > 0) {
    const files = get(sourceFiles);
    if (files.length > 0) {
      await SavePage(files[0].fullPath, number);
    }
  }
});

export const version = "0.2.1";
export default app;
