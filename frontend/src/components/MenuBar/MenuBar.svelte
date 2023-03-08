<script lang="ts">
  import { resetFolders } from "../../helpers";
  import { sourceFiles, uiChunkPage, uiChunkSize } from "../../stores";
  import ColumnSizePicker from "./ColumnSizePicker.svelte";
  import PageSelector from "./PageSelector.svelte";

  function pagePrompt() {
    // Prompt user to select a page
    const answer = prompt("Enter a page number between 1 and " + Math.floor($sourceFiles.length / $uiChunkSize));
    if (answer) {
      const page = parseInt(answer);
      if (page > 0 && page <= Math.floor($sourceFiles.length / $uiChunkSize)) {
        uiChunkPage.set(page - 1);
      }
    }
  }

</script>
<nav>
  {#if $sourceFiles.length > 0}
    <button on:click={resetFolders}>
      Start over
    </button>
    <PageSelector />
    <ColumnSizePicker />
  {/if}
  
</nav>

<style lang="scss">
  nav {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    background-color: #4B70AA;
    color: white;
    padding: 8px;

    button {
      background-color: white;
      font-weight: bold;
      color: #4B70AA;
      border-radius: 4px;
      border: none;
      padding: 5px;
      cursor: pointer;

      &:hover {
        background-color: #E0E0E0;
      }

      &:active {
        background-color: #BDBDBD;
      }

      &:focus {
        outline: none;
      }
    }
  }
</style>