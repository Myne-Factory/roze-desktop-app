
<script lang="ts">
  import { chunkPrompt, goToImagePrompt, pagePrompt } from "../../helpers";
  import { sourceFiles, uiChunkPage, uiChunkSize } from "../../stores";
  import AnimatedNumber from "../AnimatedNumber.svelte";
</script>

<div>
  <button on:click={() => uiChunkPage.update((n) => n - 1)} disabled={$uiChunkPage === 0}>
    Previous
  </button>
  <span class="page-details">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <span on:click={pagePrompt} class="page-number" title="Click to change page">
      Page {$uiChunkPage + 1} / {Math.ceil($sourceFiles.length / $uiChunkSize)} </span>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <span class="page-range" title="Click to change chunk size" on:click={chunkPrompt}>
      ({$uiChunkPage * $uiChunkSize + 1} - {Math.min($uiChunkPage * $uiChunkSize + $uiChunkSize, $sourceFiles.length)})
    </span>
  </span>
  <button on:click={() => uiChunkPage.update((n) => n + 1)} disabled={$uiChunkPage === Math.floor($sourceFiles.length / 100)}>
    Next
  </button>
  <button on:click={goToImagePrompt}>
    Go to image
  </button>
  <div class="image-count"><AnimatedNumber delay={500} duration={2500} value={$sourceFiles.length} /> <span>images</span></div>
</div>

<style lang="scss">
  div {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    color: white;

    .image-count {
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;
      margin: 0 10px;

      span {
        margin-left: 5px;
      }
    }

    button {
      background-color: white;
      font-weight: bold;
      color: #4B70AA;
      border-radius: 4px;
      border: none;
      padding: 5px;
      margin: 0 5px;
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

  span {
    cursor: pointer;

    .page-number {
      font-weight: bold;
      margin-right: 5px;
      &:hover {
        text-decoration: underline;
      }

      &:active {
        text-decoration: none;
      }

      &:focus {
        outline: none;
      }
    }

    .page-range {
      font-size: 0.8em;
      &:hover {
        text-decoration: underline;
      }

      &:active {
        text-decoration: none;
      }

      &:focus {
        outline: none;
      }
    }
  }

  .page-details {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    margin: 0 10px;
  }

</style>