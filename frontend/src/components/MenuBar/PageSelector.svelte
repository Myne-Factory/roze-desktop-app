
<script lang="ts">
  import { sourceFiles, uiChunkPage, uiChunkSize } from "../../stores";

  function goToImagePrompt() {
    const index = prompt(`Enter an image index between 1 and ${$sourceFiles.length}`, "1");

    if (!index) {
      return;
    }

    if (parseInt(index) < 1 || parseInt(index) > $sourceFiles.length) {
      alert("Index must be between 1 and " + $sourceFiles.length);
      return;
    }

    uiChunkPage.set(Math.floor((parseInt(index) - 1) / $uiChunkSize));
  }

  function pagePrompt() {
    const page = parseInt(prompt("Page number", ($uiChunkPage + 1).toString()));
    const maxPage = Math.floor($sourceFiles.length / $uiChunkSize) + 1;
    if (page > maxPage) {
      alert("Page number must be less than " + maxPage + ", setting to max page");
      uiChunkPage.set(maxPage - 1);
      return;
    }

    if (page < 1) {
      alert("Page number must be greater than 0, setting to page 1");
      uiChunkPage.set(0);
      return;
    }

    if (page) {
      uiChunkPage.set(page - 1);
    }
  };

  function chunkPrompt() {
    const chunk = prompt("Chunk size", $uiChunkSize.toString());


    if (chunk) {
      if (parseInt(chunk) < 1) {
        alert("Chunk size must be at least 1");
        return;
      }

      if (parseInt(chunk) > 500 && parseInt(chunk) < $sourceFiles.length) {
        const confirm = prompt("Are you sure you want to set the chunk size to " + chunk + "? This may cause performance issues.");
        if (confirm !== "yes") {
          uiChunkSize.set(parseInt(chunk));
          return;
        }
        return;
      }

      if (parseInt(chunk) > $sourceFiles.length) {
        alert("Setting chunk to max value of " + $sourceFiles.length);
        uiChunkSize.set($sourceFiles.length);
        return;
      }
      uiChunkSize.set(parseInt(chunk));
    }
  }
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
  <span>{$sourceFiles.length} images</span>
</div>

<style lang="scss">
  div {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    color: white;
    padding: 8px;

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