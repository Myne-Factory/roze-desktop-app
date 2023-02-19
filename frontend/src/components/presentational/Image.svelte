<script lang="ts">
  import { onMount } from "svelte";
  import { CopyFile, DeleteFile, FileExists, GetBase64Data } from "../../../wailsjs/go/main/App";
  import type { main } from "../../../wailsjs/go/models";
  import {targetFolder, uiHideCopied, fileBase64Index, type FileBase64Index, uiChunkPage} from '../../stores'

  export let file: main.Image;
  export let page: number = 0;
  let fileExistInTargetFolder: boolean = false;
  
  uiChunkPage.subscribe(value => {
    const distFromPage = Math.abs(page - value);
    if(distFromPage > 2) {
      // Clear base64 data if page is more than 1 away
      
    }
  });

  let fileBase64IndexRef:FileBase64Index = {}
  fileBase64Index.subscribe(value => {
    fileBase64IndexRef = value;
  });

  onMount(async function () {
    if(!fileBase64IndexRef[file.name]) {
      let base64 = await GetBase64Data(file.path);
      fileBase64IndexRef[file.name] = base64;
    }
    fileExistInTargetFolder = await checkIfFileExistInTargetFolder();
  });

  let targetPathRef = "";

  targetFolder.subscribe(value => {
    targetPathRef = value;
  });

  async function checkIfFileExistInTargetFolder() {
    const fileExists = await FileExists(targetPathRef, file.name);
    return fileExists;
  }

  async function handleFile() {
    if(fileExistInTargetFolder) {
      await DeleteFile(targetPathRef + "/" + file.name);
      fileExistInTargetFolder = await checkIfFileExistInTargetFolder();
    }else {
      await copyFileToTargetFolder();
    }
  }

  async function copyFileToTargetFolder() {
    fileExistInTargetFolder = await checkIfFileExistInTargetFolder();
    if(fileExistInTargetFolder) return;
    await CopyFile(file.path, targetPathRef + "/" + file.name);
    fileExistInTargetFolder = true;
  }

  let hiddenRef: boolean = false;
  uiHideCopied.subscribe(value => {
    hiddenRef = value;
  });

</script>
<!--<img src={"data:image/*;base64," + base64} alt="">-->
<div class="container" on:click={handleFile}>
  <div class="overlay" class:exists={fileExistInTargetFolder}>
    <span>✓</span>
  </div>
  {#if !fileBase64IndexRef[file.name]}
    <img src="https://logo.mynefactory.ai" alt="">
  {:else}
    <img src={"data:image/*;base64," + fileBase64IndexRef[file.name]} alt="">
  {/if}
</div>
<style>
  .overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    font-size: 5rem;
    /* Strong text shadow */
    text-shadow: 0 0 1px rgba(0, 0, 0, 0.5), 0 0 2px rgba(0, 0, 0, 0.5), 0 0 3px rgba(0, 0, 0, 0.5), 0 0 4px rgba(0, 0, 0, 0.5), 0 0 5px rgba(0, 0, 0, 0.5), 0 0 6px rgba(0, 0, 0, 0.5);

    transition: opacity 0.3s ease-in-out;
    opacity: 0;
  }

  .overlay.exists {
    opacity: 1;
  }

  .container {
    position: relative;
    width: 100%;
    height: 100%;
    user-select: none;
    cursor: pointer;
  }
  

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    pointer-events: none;
  }
</style>