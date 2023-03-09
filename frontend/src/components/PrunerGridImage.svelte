<script lang="ts">
  import { targetFolder } from '../../src/stores';
  import type { main } from 'wailsjs/go/models';
  import { CopyFile, DeleteFile, FileExists } from "../../wailsjs/go/main/App";
  import { onMount, afterUpdate } from 'svelte';
  import LogRocket from 'logrocket';

  export let file: main.Image;
  
  let fileExistInTargetFolder = false;

  onMount(async function () {
    fileExistInTargetFolder = await checkIfFileExistInTargetFolder();
  });

  afterUpdate(async function () {
    fileExistInTargetFolder = await checkIfFileExistInTargetFolder();
  });

  async function checkIfFileExistInTargetFolder() {
    const fileExists = await FileExists($targetFolder, file.name);
    return fileExists;
  }

  async function handleFile() {
    if(fileExistInTargetFolder) {
      await DeleteFile($targetFolder + "/" + file.name);
      fileExistInTargetFolder = await checkIfFileExistInTargetFolder();
    }else {
      await copyFileToTargetFolder();
    }
  }

  async function copyFileToTargetFolder() {
    fileExistInTargetFolder = await checkIfFileExistInTargetFolder();
    if(fileExistInTargetFolder) return;
    await CopyFile(file.path, $targetFolder + "/" + file.name);
    fileExistInTargetFolder = true;
  }

</script>

<div class="gallery-item" on:click={handleFile}>
  <img src="" data-private
    data-drive={file.drive} 
    data-path={file.path} 
    alt={file.name} 
    style="opacity: 0" />
  {#if fileExistInTargetFolder}
    <div class="copied">Copied</div>
  {/if}
  <div class="overlay">
    <div class="file-tip">Click to {file.copied ? 're-copy' : 'copy'}</div>
  </div>
  <div class="loader"></div>
</div>

<style lang="scss">
  .gallery-item {
    position: relative;
    padding: 2.5px;
    min-height: 200px;
    transition: min-height 0.3s ease-in-out;
    cursor: pointer;

    img {
      display: block;
      width: 100%;
      height: auto;
      transition: opacity 0.3s ease-in-out;
    }

    /**Always visible*/
    .copied {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background-color: rgba(0, 0, 0, 0.5);
      z-index: 1;
      opacity: 1;
      transition: opacity 0.3s ease-in-out;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 1.2rem;
      font-weight: bold;
      color: white;
    }


    .overlay {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background-color: rgba(0, 0, 0, 0.5);
      z-index: 1;
      opacity: 0;
      transition: opacity 0.3s ease-in-out;


      .file-tip {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        color: white;
        text-shadow: 0 0 5px black;
        background-color: rgba(0, 0, 0, 0.5);
        font-size: 0.8rem;
      }
    }

    &:hover .overlay {
      opacity: 1;
    }

    .loader {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 40px;
      height: 40px;
      border: 3px solid rgba(255, 255, 255, 0.9);
      border-top-color: transparent;
      border-radius: 50%;
      animation: spin 1s linear infinite;
      display: none;
      z-index: 2;
      transition-delay: 0.1s;
    }
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>