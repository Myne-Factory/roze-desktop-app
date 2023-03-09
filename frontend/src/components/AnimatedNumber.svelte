<!--
  Animates a number from 0 to the given value.
  1 -> 100, 0.1ms between each step
-->
<script lang="ts">
  import { cubicOut } from "svelte/easing";
  import { tweened } from "svelte/motion";
  import { onMount } from "svelte";

  // The value to animate to
  export let value = 0;
  // Duration of the animation in ms
  export let duration = 1000;
  // Delay before the animation starts 
  export let delay = 0;

  let tweenedValue;

  onMount(() => {
    tweenedValue = tweened(0, {
      duration,
      delay,
      easing: cubicOut,
    });
    tweenedValue.set(value);
  });

  $: if (tweenedValue) {
    tweenedValue.set(value);
  }

  function formatNumber(number) {
    return Math.round(number);
  }
</script>

<div class="number">
  {#if tweenedValue}
    {formatNumber($tweenedValue)}
  {:else}
    {formatNumber(value)}
  {/if}
</div>

<style lang="scss">
  .number {
    font-size: 0.8rem;
    font-weight: bold;
    text-align: center;
  }
</style>