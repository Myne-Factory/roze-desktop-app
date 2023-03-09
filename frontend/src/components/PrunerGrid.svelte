<script lang="ts">
  import { sourceFiles, uiColumns, uiChunkPage, uiChunkSize, loading, uiChunkPageScrollMemory } from "../stores";
  import { afterUpdate, onMount } from "svelte";
  import PrunerGridImage from "./PrunerGridImage.svelte";
  import { initLogRocket } from "../helpers";

  let galleryWidth = 0;
  let gap = 5;
  let maxColumnWidth = 0;
  
  $: maxColumnWidth = galleryWidth / $uiColumns;
  $: galleryStyle = `grid-template-columns: repeat(auto-fit, minmax(0, ${maxColumnWidth}px));`;

  function createObserver() {
    const observer = new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const element = entry.target as HTMLDivElement;
          const img = element.querySelector("img") as HTMLImageElement;
          const loader = element.querySelector(".loader") as HTMLDivElement;
          loader.style.display = "block";
          img.onload = () => {
            loader.style.display = "none";
            img.style.opacity = "1";
            img.style.transition = "min-height 0.3s ease-in-out"; // add transition to min-height

            // Set height to img height
            element.style.height = `${img.height}px`;
            element.style.minHeight = `${img.height}px`;
            observer.unobserve(entry.target);
          };
          img.src = `/files${img.dataset.drive}:${img.dataset.path}`;
          img.style.opacity = "0";
          img.removeAttribute("loading");
        }
      });
    });
    
    document.querySelectorAll(".gallery-item").forEach((item) => {
      observer.observe(item);
    });
    
    return observer;
  }

  let observer = createObserver();

  export function recreateObserver() {
    observer.disconnect();
    observer = createObserver();
  }


  uiChunkPage.subscribe(() => {
    recreateObserver();
  });

  uiChunkSize.subscribe(() => {
    recreateObserver();
  });
  
  afterUpdate(() => {
    recreateObserver();
  });

  function saveScrollMemory() {
    const container = this as HTMLDivElement;
    const key = $uiChunkPage;
    const value = container.scrollTop;
    uiChunkPageScrollMemory.update((memory) => {
      memory[key] = value;
      return memory;
    });
  }

 
</script>

<div class="container" bind:clientWidth={galleryWidth} on:scroll={saveScrollMemory}>
  <div class="gallery" style={galleryStyle}>
    {#each $sourceFiles.slice($uiChunkPage * $uiChunkSize, ($uiChunkPage + 1) * $uiChunkSize) as file}
      <PrunerGridImage {file} />
    {/each}
  </div>
</div>

<style lang="scss">
  .container {
    overflow-y: auto; 
  }

  .gallery {
    display: grid;
    grid-auto-flow: row dense;
  }
</style>