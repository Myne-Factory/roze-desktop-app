<script lang="ts">
</script>

<!--<div class="logrocket-container">
  <div class="logrocket-icon"></div>
  <label class="logrocket-label" on:click={enableLogging} class:logrocket-active={$logrocketRecording}>
    <div class="logrocket-text">{$logrocketRecording ? "Session recording is enabled. Thanks for helping us improve the app!" : "Enable session recording to help us improve the app"}</div>
    <div class="logrocket-text"><b>{$logrocketRecording ? "Close this window to stop recording" : ""}</b></div>
  </label>
</div>-->
<style lang="scss">
  .logrocket-container {
    display: flex;
    align-items: center;
    margin-top: 20px;
    padding: 20px;
    background-color: white;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    transition: all 0.2s ease-in-out;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.15);
    }
  }

  .logrocket-icon {
    width: 20px;
    height: 20px;
    background-image: url('logrocket-icon.png');
    background-size: contain;
    background-repeat: no-repeat;
    margin-right: 10px;
  }

  .logrocket-label {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 0;
    cursor: pointer;
    position: relative;

    &.logrocket-active {
      .logrocket-text {
        color: #44a46a;
        font-weight: bold;
      }

    }

    input[type="checkbox"] {
      position: absolute;
      opacity: 0;

      &:checked + .logrocket-text {
        color: #44a46a;
        font-weight: bold;
      }

      &:checked ~ .logrocket-icon {
        background-image: url('logrocket-icon-active.png');
      }
    }
  }

  .logrocket-text {
    font-size: 18px;
    color: #858585;
    transition: color 0.2s ease-in-out;
  }
</style>