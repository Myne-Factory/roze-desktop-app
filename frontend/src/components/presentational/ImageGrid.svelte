<script lang="ts">
  import { sourceFiles, uiChunkPage, uiChunkSize, uiColumns } from "../../../src/stores";
  import Image from "./Image.svelte";
  import { onMount } from "svelte";

  onMount(() => {
    // Move 50 pixels down to hide the top of the first image
    document.querySelector(".wrapper").scrollTop = 50;
  });

  let columnRef: number = 5;

  uiColumns.subscribe(value => {
    columnRef = value;
  });

  let sourceFilesRef = [];

  sourceFiles.subscribe(value => {
    sourceFilesRef = value;
  });

  let chunkPageRef = 0;
  uiChunkPage.subscribe(value => {
    chunkPageRef = value;
  });

  let chunkSizeRef = 400;
  uiChunkSize.subscribe(value => {
    chunkSizeRef = value;
  });

  // If the user scrolls to the bottom, load the next chunk page, if possible
  // if the user scrolls to the top, load the previous chunk page, if possible
  function onScroll(e) {
    const { scrollTop, scrollHeight, clientHeight } = e.target;

    if (scrollTop + clientHeight >= scrollHeight && chunkPageRef < (sourceFilesRef.length / chunkSizeRef) - 1) {
      uiChunkPage.update(value => value + 1);
      // scroll top +50 to hide the top of the first image
      e.target.scrollTop = 50;
    } else if (scrollTop <= 0 && chunkPageRef > 0) {
      uiChunkPage.update(value => value - 1);
      // scroll bottom -50 to hide the bottom of the last image
      e.target.scrollTop = scrollHeight - clientHeight - 50;
      
    }
  }
</script>

<div class="wrapper" on:scroll={onScroll}>
  <div class="image-grid"  style="grid-template-columns: repeat({columnRef}, 1fr)">
    <!-- Load only chunkSizeRef images at a time, offset by chunkPageRef -->
    {#each sourceFilesRef.slice(chunkPageRef * chunkSizeRef, (chunkPageRef + 1) * chunkSizeRef) as file (file.path)}
      <Image file={file} page={chunkPageRef} />
    {/each}
  </div>
</div>

<style>
  .wrapper {
    padding-top: 50px;
    padding-bottom: 50px;
    flex-grow: 1;
    overflow-y: auto; 
  }

  /**
   * 5 columns
  */
  .image-grid {
    display: grid;
    grid-gap: 4px;
    padding: 4px;
  }
</style>