<script>
  import { sourceFiles, uiColumns, uiChunkPage, uiChunkSize, loading } from "../stores";

  let galleryWidth = 0;
  let gap = 5;
  
  $: maxColumnWidth = galleryWidth/($uiColumns);

  $: galleryStyle = `grid-template-columns: repeat(auto-fit, minmax(0, ${maxColumnWidth}px)); grid-gap: ${gap}px`;
</script>

<style lang="scss">
  .gallery {
    display: grid;
    grid-gap: 5px;
  }

</style>

<div class="gallery" bind:clientWidth={galleryWidth} style={galleryStyle}>
  <div>
    <div class="overlay">
      Image
    </div>
    <slot></slot>
  </div>
</div>

