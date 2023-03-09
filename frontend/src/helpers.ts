import { get } from "svelte/store";
import { DiagFolderPath, ListImages } from "../wailsjs/go/main/App";
import {
  sourceFolder,
  targetFolder,
  sourceFiles,
  loading,
  uiChunkPage,
  uiChunkSize,
  logrocketRecording,
} from "./stores";

import LogRocket from "logrocket";

export async function openDirectoryDialog(title: string) {
  return await DiagFolderPath(title);
}

export async function loadSourceFolder() {
  const path = await openDirectoryDialog("Select pruning/source folder");
  if (!path) return;
  sourceFolder.set(path);

  // Fetch the source files
  await fetchFiles();
}

export async function loadTargetFolder() {
  const path = await openDirectoryDialog("Select target/output folder");
  if (!path) return;
  targetFolder.set(path);

  // Fetch the source files
  await fetchFiles();
}

export async function resetFolders() {
  sourceFolder.set(null);
  targetFolder.set(null);
  sourceFiles.set([]);
}

export async function fetchFiles() {
  const src = get(sourceFolder);
  const tgt = get(targetFolder);
  // If source and target folder are set, fetch the files
  // and store them in the sourceFiles store
  if (!src || !tgt) {
    return;
  }

  loading.set(true);
  const files = await ListImages(src, tgt);
  console.log(files);
  console.log("Length of sourceFiles: ", files.length);
  sourceFiles.set(files);
  loading.set(false);
}

export function goToImagePrompt() {
  const index = prompt(
    `Enter an image index between 1 and ${get(sourceFiles).length}`,
    "1"
  );

  if (!index) {
    return;
  }

  if (parseInt(index) < 1 || parseInt(index) > get(sourceFiles).length) {
    alert("Index must be between 1 and " + get(sourceFiles).length);
    return;
  }

  uiChunkPage.set(Math.floor((parseInt(index) - 1) / get(uiChunkSize)));
}

export function pagePrompt() {
  const page = parseInt(
    prompt("Page number", (get(uiChunkPage) + 1).toString())
  );
  const maxPage = Math.floor(get(sourceFiles).length / get(uiChunkSize)) + 1;
  if (page > maxPage) {
    alert("Page number must be less than " + maxPage + ", setting to max page");
    uiChunkPage.set(maxPage - 1);
    return;
  }

  if (page < 1) {
    alert("Page number must be greater than 0, setting to page 1");
    uiChunkPage.set(0);
    return;
  }

  if (page) {
    uiChunkPage.set(page - 1);
  }
}

export function chunkPrompt() {
  const chunk = prompt("Chunk size", get(uiChunkSize).toString());

  if (chunk) {
    if (parseInt(chunk) < 1) {
      alert("Chunk size must be at least 1");
      return;
    }

    if (parseInt(chunk) > 500 && parseInt(chunk) < get(sourceFiles).length) {
      const conf = confirm(
        "Are you sure you want to set the chunk size to " +
          chunk +
          "? This may cause performance issues."
      );
      if (conf) {
        uiChunkSize.set(parseInt(chunk));
        return;
      }
      return;
    }

    if (parseInt(chunk) > get(sourceFiles).length) {
      alert("Setting chunk to max value of " + get(sourceFiles).length);
      uiChunkSize.set(get(sourceFiles).length);
      return;
    }
    uiChunkSize.set(parseInt(chunk));
  }
}
