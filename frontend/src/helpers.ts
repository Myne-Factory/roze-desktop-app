import {
  DiagFolderPath,
  ListImages,
  GetBase64Data,
} from "../wailsjs/go/main/App";
import { sourceFolder, targetFolder, sourceFiles } from "./stores";

export async function openDirectoryDialog(title: string) {
  return await DiagFolderPath(title);
}

export async function loadSourceFolder() {
  const path = await openDirectoryDialog("Select pruning/source folder");
  sourceFolder.set(path);
  fetchSourceFiles(path);
}

export async function loadTargetFolder() {
  const path = await openDirectoryDialog("Select target/output folder");
  targetFolder.set(path);
}

export async function fetchSourceFiles(path: string) {
  // If source and target folder are set, fetch the files
  // and store them in the sourceFiles store

  if (sourceFolder && targetFolder) {
    const files = await ListImages(path);
    console.log(files);
    console.log("Length of sourceFiles: ", files.length);
    sourceFiles.set(files);
  }
}
