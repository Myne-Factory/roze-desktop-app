<script lang="ts">
  import Chat from "./components/chatgpt/Chat.svelte";
  import Logo from "./components/Logo.svelte";
  import MenuBar from "./components/MenuBar/MenuBar.svelte";
  import PrunerGrid from "./components/PrunerGrid.svelte";

  import { loadSourceFolder, loadTargetFolder } from "./helpers";
  import { sourceFolder, targetFolder } from "./stores"
  // If both folders are set, set true
  $: isReady = $sourceFolder && $targetFolder
  
</script>

<main>
  <MenuBar />
  {#if isReady}
    <PrunerGrid />
  {:else}
    <div class="not-ready">
      <Logo />
      <div>Choose source and target folder.</div>
        <div class="folder-wrapper">
          <div class="folder-selectors">
            <button on:click={loadSourceFolder}>Load Source</button>
            <span>{$sourceFolder || 'No source selected'}</span>
          </div>
          <div class="folder-selectors">
            <button on:click={loadTargetFolder}>Load Target</button>
            <span>{$targetFolder || 'No target selected'}</span>
          </div>
        </div>
    </div>
  {/if}
</main>
 
<style lang="scss">
  main {
    display: flex;
    flex-direction: column;
    color: black;
    height: 100vh;
    .not-ready {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      height: 100%;

      .folder-wrapper {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: space-around;
      margin-top: 20px;

        .folder-selectors {
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          height: 60px;
          margin: 10px 0px;

          span {
            flex-grow: 1;
            padding: 10px;
          }
        }
      }
    }
  }

    button {
      background-color: #616161;
      border: none;
      color: white;
      padding: 5px;
      font-size: 16px;
      cursor: pointer;

      &:hover {
        background-color: #555;
      }

      &:active {
        background-color: #777;
      }

      &:focus {
        outline: none;
      }
    }
</style>
