<script lang="ts">
  import { sourceFiles, uiColumns, uiChunkPage, uiChunkSize, loading } from "../stores";
  import { onMount } from "svelte";

  let galleryWidth = 0;
  let gap = 5;
  let observer: IntersectionObserver;
  
  $: maxColumnWidth = galleryWidth/($uiColumns);
  $: galleryStyle = `grid-template-columns: repeat(auto-fit, minmax(0, ${maxColumnWidth}px));`;

  onMount(() => {
    observer = new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const img = entry.target.querySelector("img") as HTMLImageElement;
          const loader = entry.target.querySelector(".loader") as HTMLDivElement;
          loader.style.display = "block";
          img.onload = () => {
            loader.style.display = "none";
            img.style.opacity = "1";
            img.style.transition = "min-height 0.3s ease-in-out"; // add transition to min-height
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
  });
</script>

<div class="container" bind:clientWidth={galleryWidth}>
  <div class="loading-overlay" style="display: {$loading ? 'block' : 'none'}">
    <div class="loading-spinner"></div>
  </div>
  <div class="gallery" style={galleryStyle}>
    {#each $sourceFiles.slice($uiChunkPage * $uiChunkSize, ($uiChunkPage + 1) * $uiChunkSize) as file (file.path)}
      <div class="gallery-item">
        <img src="" 
             data-drive={file.drive} 
             data-path={file.path} 
             alt={file.name} 
             style="opacity: 0" />
        <div class="overlay">
          <div class="file-details">
            <div class="file-name">{file.name}</div>
          </div>
        </div>
        <div class="loader"></div>
      </div>
    {/each}
  </div>
</div>

<style lang="scss">
  .container {
    overflow-y: auto; 
  }

  .loading-overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 100;
  }

  .gallery {
    display: grid;
    grid-auto-flow: dense;
  }

  .gallery-item {
    position: relative;
    padding:2.5px;
  }

  .gallery-item img {
    display: block;
    width: 100%;
    height: auto;
    transition: opacity 0.3s ease-in-out;
  }

  .gallery-item .overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 1;
    opacity: 0;
    transition: opacity 0.3s ease-in-out;

    &:hover .overlay {
      opacity: 1;

      .file-details {
        transform: translateY(0);
        transition-delay: 0.2s;
      }
    }

    .file-details {
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
      padding: 5px;
      color: white;
      font-size: 0.8rem;
      background-color: rgba(0, 0, 0, 0.5);
      text-align: center;
      transform: translateY(100%);
      transition: transform 0.2s ease-out;
    }
  }

  .gallery-item .overlay,
  .gallery-item .loader {
    transition-delay: 0.1s;
  }

  .gallery-item:hover .overlay {
    opacity: 1;
  }

  .gallery-item .loader {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 40px;
    height: 40px;
    border: 3px solid rgba(255, 255, 255, 0.9);
    border-top-color: transparent;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    display: none;
    z-index: 2;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>


