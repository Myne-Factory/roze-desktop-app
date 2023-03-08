import { get } from "svelte/store";
import { DiagFolderPath, ListImages } from "../wailsjs/go/main/App";
import { sourceFolder, targetFolder, sourceFiles, loading } from "./stores";

export async function openDirectoryDialog(title: string) {
  return await DiagFolderPath(title);
}

export async function loadSourceFolder() {
  const path = await openDirectoryDialog("Select pruning/source folder");
  if (!path) return;
  sourceFolder.set(path);

  // Fetch the source files
  fetchSourceFiles(get(sourceFolder));
}

export async function loadTargetFolder() {
  const path = await openDirectoryDialog("Select target/output folder");
  if (!path) return;
  targetFolder.set(path);

  // Fetch the source files
  fetchSourceFiles(get(sourceFolder));
}

export async function resetFolders() {
  sourceFolder.set(null);
  targetFolder.set(null);
  sourceFiles.set([]);
}

export async function fetchSourceFiles(path: string) {
  // If source and target folder are set, fetch the files
  // and store them in the sourceFiles store
  if (!sourceFolder || !targetFolder) {
    return;
  }

  loading.set(true);
  const files = await ListImages(path);
  console.log(files);
  console.log("Length of sourceFiles: ", files.length);
  sourceFiles.set(files);
  loading.set(false);
}
