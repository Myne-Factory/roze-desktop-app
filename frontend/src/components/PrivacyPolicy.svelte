<script>
  import { showPrivacyPolicy } from "../stores";

  export let lastUpdated = '2023-03-9';
  export let contactEmail = 'mail@mynefactory.ai';

  function handleClose() {
    showPrivacyPolicy.set(false);
  }
</script>

{#if $showPrivacyPolicy}
  <div class="overlay">
    <div class="privacy-policy">
      <button class="close-button" on:click={handleClose}>X</button>
      <h2>Privacy Policy</h2>
      <p>Last updated: {lastUpdated}</p>
      <p>We take your privacy seriously. This privacy policy describes how we collect, use, and share information when you use <span class="important">Roze</span>.</p>
      <h3>Information We Collect</h3>
      <p>When you use our app, we may collect the following types of information:</p>
      <ul>
        <li>App usage data: We may collect information about how you use the app, such as your interactions with its features, error messages, and other usage statistics.</li>
        <li>App interface images: If you opt-in to session recording, we may capture images of the app interface for the purpose of improving the app's performance and usability. We will not capture any personal images or file paths stored on your computer.</li>
      </ul>
      <h3>How We Use Your Information</h3>
      <p>We may use the information we collect to:</p>
      <ul>
        <li>Improve the app's performance and usability.</li>
        <li>Identify and troubleshoot technical issues.</li>
        <li>Respond to your support requests.</li>
      </ul>
      <h3>Information Sharing</h3>
      <p>We will not share your personal information with third parties unless:</p>
      <ul>
        <li>We have your explicit consent to do so.</li>
        <li>We are required to do so by law.</li>
      </ul>
      <h3>Security</h3>
      <p>We take reasonable measures to protect your information from unauthorized access, disclosure, alteration, or destruction.</p>
      <h3>Changes to this Policy</h3>
      <p>We may update this privacy policy from time to time. We will notify you of any changes by posting the new policy on our website.</p>
      <h3>Contact Us</h3>
      <p>If you have any questions or concerns about this privacy policy, please contact us at {contactEmail}.</p>
    </div>
  </div>
{/if}

<style>
  .important {
    font-weight: bold;
  }

  .overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 999;
  }

  .privacy-policy {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: white;
    padding: 20px;
    width: 80%;
    max-width: 600px;
    border-radius: 10px;
    box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.3);
    text-align: left;
    font-size: 16px;
    overflow-y: scroll;
    max-height: calc(100% - 100px);
    padding-top: 50px;
  }

  .close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  border: none;
  background-color: white;
  font-size: 24px;
  cursor: pointer;
  outline: none;
  }

  h2 {
  margin-top: 0;
  font-size: 24px;
  }

  h3 {
  margin-top: 20px;
  margin-bottom: 10px;
  font-size: 18px;
  }

  ul {
  list-style: disc;
  margin-top: 10px;
  margin-bottom: 10px;
  }

  li {
  margin-left: 20px;
}
</style>