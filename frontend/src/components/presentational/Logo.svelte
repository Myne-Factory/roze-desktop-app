<!-- Logo.svelte -->
<script lang="ts">
  import { loadSourceFolder, loadTargetFolder } from "../../../src/helpers";

  import {sourceFolder, targetFolder, uiChunkSize} from "../../stores"
  
  let sourceFolderRef: string
  let targetFolderRef:string

  sourceFolder.subscribe(value => {
    sourceFolderRef = value
  })

  targetFolder.subscribe(value => {
    targetFolderRef = value
  })

  let chunkSizeRef = 1
  uiChunkSize.subscribe(value => {
    chunkSizeRef = value
  })
</script>

<div class="container">
  <img src="https://logo.mynefactory.ai" alt="">
  
  <div class="buttons">
    {#if sourceFolderRef}
    <div class="ready">
      <h3>Source folder</h3>
      <p>{sourceFolderRef}</p>
      <span>Good, now select a target folder</span>
    </div>
  {:else}
    <button class="fancy-button" on:click={loadSourceFolder}>
      Select a folder to import
      <span>The folder you'll prune</span>
    </button>
  {/if}

  {#if targetFolderRef}
    <div class="ready">
      <h3>Target folder</h3>
      <p>{targetFolderRef}</p>
      <span>Good, now select a source folder</span>
    </div>
  {:else}
    <button class="fancy-button" on:click={loadTargetFolder}>
      Select a folder to export
      <span>The folder you'll prune</span>
    </button>
  {/if}
  </div>
</div>

<style>
  div.container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 5rem;
    flex: 1;
  }

  div.buttons {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    margin-top: 2rem;
  }

  div.option {
    background-color: #f5f5f5;
    color: #333;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-top: 2rem;
  }

  div.ready {
    color: black;
    text-align: center;
    margin: 1rem;
  }

  div.ready > h3 {
    font-size: 1.5rem;
    padding: 0;
    margin: 0;
  }

  div.ready > p {
    padding: 4px;
    margin: 0;
  }

  button > span {
    display: block;
    font-size: 0.8rem;
  }

  img {
    max-width: 100%;
    max-height: 100%;
  }
</style>