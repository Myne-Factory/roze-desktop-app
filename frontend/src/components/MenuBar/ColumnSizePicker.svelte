<script lang="ts">
  import { uiColumns } from "../../stores";

  let minColumns = 1;
  let maxColumns = 10;
</script>

<div class="column-picker">
  <button on:click={() => uiColumns.set($uiColumns - 1)} disabled={$uiColumns <= minColumns}>
    - 
  </button>
  <span>Columns: {$uiColumns}</span>
  <button on:click={() => uiColumns.set($uiColumns + 1)} disabled={$uiColumns >= maxColumns}>
    +
  </button>
</div>

<style lang="scss">
  .column-picker {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    button {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 30px;
      height: 30px;
      border-radius: 4px;
      border: 1px solid black;
      background-color: white;
      font-weight: bold;
      font-size: 1.2em;
      border: none;
      cursor: pointer;
      &:disabled {
        background-color: lightgray;
        cursor: not-allowed;
      }
    }
    span {
      margin: 0px 10px;
    }
  }
</style>