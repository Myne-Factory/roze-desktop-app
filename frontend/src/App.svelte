<script lang="ts">
  import MenuBar from "./components/presentational/MenuBar.svelte";
  import Logo from "./components/presentational/Logo.svelte";

  import {sourceFolder, targetFolder} from "./stores"
  import ImageGrid from "./components/presentational/ImageGrid.svelte";
  
  let sourceFolderRef
  let targetFolderRef

  sourceFolder.subscribe(value => {
    sourceFolderRef = value
  })

  targetFolder.subscribe(value => {
    targetFolderRef = value
  })

  // If both folders are set, set true
  $: isReady = sourceFolderRef && targetFolderRef
  
</script>

<main>
  <MenuBar />
  {#if isReady}
    <ImageGrid />
  {:else}
    <Logo />
  {/if}

</main>
 
<style>
  main {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }
</style>
