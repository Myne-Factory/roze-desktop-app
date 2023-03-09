<script lang="ts">
  import CheckIcon from "./components/CheckIcon.svelte";
  import Logo from "./components/Logo.svelte";
  import MenuBar from "./components/MenuBar/MenuBar.svelte";
  import PrunerGrid from "./components/PrunerGrid.svelte";
  import PrivacyPolicy from "./components/PrivacyPolicy.svelte";

  import { loadSourceFolder, loadTargetFolder } from "./helpers";
  import { isReady, showPrivacyPolicy, sourceFiles, sourceFolder, targetFolder } from "./stores"
  import LogrocketOptin from "./components/LogrocketOptin.svelte";
  // If both folders are set, set true
  

  async function checkIfReady() {
    if($sourceFolder && $targetFolder) {
      setTimeout(() => {
        isReady.set(true)
      }, 500);
    }
  }

  async function loadSource() {
    await loadSourceFolder()
    await checkIfReady()
  }

  async function loadTarget() {
    await loadTargetFolder()
    await checkIfReady()
  }
</script>

<main>
  <PrivacyPolicy />
  {#if $isReady}
    {#if $sourceFiles.length > 0}
    <MenuBar />
    {/if}
    <PrunerGrid />
  {:else}
    <div class="not-ready">
      <Logo />
      {#if $sourceFolder && $targetFolder}
        <h3>Let's get started!</h3>
      {:else}
        <h3>Choose source and target folder.</h3>
      {/if}
        <div class="folder-wrapper">
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <div class="folder-selectors" on:click={loadSource} class:is-ready={$sourceFolder} class:transition-out={$sourceFolder && $targetFolder}>
            <span class="folder-type">Load Source</span>
            <span class="folder-help" data-private>{$sourceFolder || 'No source selected'}</span>
            {#if $sourceFolder}
              <CheckIcon />
              
            {/if}
          </div>
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <div class="folder-selectors"on:click={loadTarget} class:is-ready={$targetFolder} class:transition-out={$sourceFolder && $targetFolder}>
            <span class="folder-type" >Load Target</span>
            <span class="folder-help" data-private>{$targetFolder || 'No target selected'}</span>
            {#if $targetFolder}
              <CheckIcon />
            {/if}
          </div>
        </div>
        <LogrocketOptin />
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
        flex-direction: row;
        align-items: center;
        justify-content: space-around;
        margin-top: 20px;
        

        .folder-selectors {
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          margin: 10px 20px;
          border: 2px solid #858585;
          background-color: #f0f0f0;
          border-radius: 8px;
          padding: 20px;
          /**Inner shadow blur 10, spread 5, opacity 10%*/
          box-shadow: inset 0 0 10px 5px rgba(0, 0, 0, 0.1);

          transition: all 0.2s ease-in-out;
          cursor: pointer;

          &:hover {
            background-color: #e0e0e0;
          }

          &:active {
            background-color: #bdbdbd;
          }
          

          .folder-type {
            font-size: 28px;
            font-weight: bold;
            margin-bottom: 2px;
          }

          .folder-help {
            font-size: 20px;
            font-style: italic;
          }
        }

        .folder-selectors.is-ready {
          color: white;
          background-color: #44A46A;

          &:hover {
            background-color: #3B8E5F;
          }

          &:active {
            background-color: #2F6F4B;
          }
        }

        .folder-selectors.transition-out {
          animation: fancy-exit 0.5s ease-in-out forwards;
        }

        .folder-selectors.transition-out:nth-child(1) {
          animation-delay: 0.2s;
        }

        .folder-selectors.transition-out:nth-child(2) {
          animation-delay: 0.4s;
        }

        @keyframes fancy-exit {
          0% {
            transform: scale(1);
          }
          50% {
            transform: scale(1);
          }
          70% {
            transform: scale(1.2);
          }
          100% {
            transform: scale(0);
          }
        }
      }
    }
  }

    
</style>
