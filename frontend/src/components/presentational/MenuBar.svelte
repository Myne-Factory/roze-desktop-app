<script lang="ts">
  import { loadTargetFolder, loadSourceFolder } from "../../../src/helpers";
  import Fa from 'svelte-fa'
  import { faFileArrowDown, faFileArrowUp, faArrowLeft, faArrowRight } from '@fortawesome/free-solid-svg-icons'

  import {sourceFolder, targetFolder, sourceFiles, uiColumns, uiHideCopied, uiChunkPage, uiChunkSize} from "../../stores"
  
  let sourceFolderRef: string
  let targetFolderRef:string

  sourceFolder.subscribe(value => {
    sourceFolderRef = value
  })

  targetFolder.subscribe(value => {
    targetFolderRef = value
  })

  let sourceFilesRef = []
  sourceFiles.subscribe(value => {
    sourceFilesRef = value
  })


  let value:number = 5
  uiColumns.subscribe(value => {
    value = value
  })

  function onColumnChange(e) {
    uiColumns.set(e.target.value)
  }

  let hiddenRef = false
  uiHideCopied.subscribe(value => {
    hiddenRef = value
  })

  function onHiddenChange(e) {
    uiHideCopied.set(e.target.checked)
  }

  let chunkPageRef = 0
  let chunkSizeRef = 1

  uiChunkPage.subscribe(value => {
    chunkPageRef = value
  })

  uiChunkSize.subscribe(value => {
    chunkSizeRef = value
  })

</script>

<nav class="mainnav"> 
  <nav class="subnav" >
  
    <div>
      <div>
        <button on:click={() => uiChunkPage.set(chunkPageRef - 1)} disabled={chunkPageRef === 0}>
          <Fa icon={faArrowLeft}/>
        </button>
        <span class="small">Page {chunkPageRef + 1} / {Math.floor(sourceFilesRef.length / chunkSizeRef) + 1}</span>
        <button on:click={() => uiChunkPage.set(chunkPageRef + 1)} disabled={chunkPageRef === Math.floor(sourceFilesRef.length / chunkSizeRef)}>
          <Fa icon={faArrowRight}/>
        </button>
      </div>
      <span class="small">
        Images: {sourceFilesRef.length}
      </span>
    </div>
  </nav>
  <div class="selected-folders">
    {sourceFolderRef || "Select source folder"}
    ->
    {targetFolderRef || "Select target folder"}
  </div>
  <div>
    <input type="range" min="1" max="10" bind:value on:change={onColumnChange} class="column-slider">
    <div>Columns: {value}</div>
  </div>
</nav>

<style>
  nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .mainnav {
    background-color: #A5A5A5;
    padding: 8px;
    background-color: #878787;
    font-size:16px;
  }

  .small {
    font-size: 12px;
  }

  .folder {
    font-size: 32px;
  }

  .selected-folders {
    flex-grow: 1;
    font-size: 16px;;
  }

  .subnav {
    
    justify-content: start;
  }

  .subnav > div {
    padding: 0 8px;
    cursor: pointer;
  }

  .column-slider {
    
  }

  button {
    /*Minimalistic button*/
    background-color: #A5A5A5;
    border: none;
    color: white;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    cursor: pointer;
  }
</style>