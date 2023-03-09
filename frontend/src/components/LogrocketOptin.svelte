<script lang="ts">
  import { initLogRocket } from "../helpers";
  import { logrocketRecording, showPrivacyPolicy } from "../stores";

  const disclosureText =  "We use LogRocket to record user sessions, which helps us identify and fix issues to improve your experience. " +
                          "When you enable session recording, the app will mask any personal information(images, file paths, text, etc) before sending it to LogRocket."

  logrocketRecording.subscribe(function() {
    initLogRocket();
  })
</script>

<div class="logrocket-container">
  <div class="logrocket-icon"></div>
  <label class="logrocket-label" on:click={() => logrocketRecording.set(true)} class:logrocket-active={$logrocketRecording}>
    <div class="logrocket-text">{$logrocketRecording ? "Session recording is enabled. Thanks for helping us improve the app!" : "Enable session recording to help us improve the app"}</div>
    <div class="logrocket-text"><b>{$logrocketRecording ? "Close this window to stop recording" : ""}</b></div>
  </label>
</div>
<span class="disclosure">{disclosureText} By enabling session recording, you agree to our <a on:click={() => showPrivacyPolicy.set(true)}>privacy policy</a>.</span>
 
<style lang="scss">

a {
  color: #858585;
  text-decoration: underline;
  cursor: pointer;
  animation: pulse 1s infinite;
}

@keyframes pulse {
  // Add a pulse animation to the link, colourful and fun! So no one misses it.
  0% {
    color: #858585;
  }

  50% {
    color: #2f80ed;
  }

  100% {
    color: #858585;
  }
}

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

  .disclosure {
      font-size: 10px;
      color: #858585;
      margin-top: 10px;
      margin-left: 10px;
      line-height: 1.5;
      max-width: 500px;
      a {
        color: #858585;
        text-decoration: underline;
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