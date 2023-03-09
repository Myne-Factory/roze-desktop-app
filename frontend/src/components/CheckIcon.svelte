<div class="checkmark">
  <svg viewBox="0 0 24 24">
    <path d="M5.5 12.5L9 16L20 5" stroke="#44A46A" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" fill="none" />
  </svg>
</div>

<style lang="scss">
  .checkmark {
    display: inline-block;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: white;
    font-size: 20px;
    line-height: 40px;
    text-align: center;
    overflow: hidden;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    animation: slide-in 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards;

    svg {
      margin-right:2px;
      margin-top: 4px;
    }
  }

  /**
   * Slide in animation
   * 0.1s of nothing
   * 0.3s of height and opacity
  */
  @keyframes slide-in {
    0% {
      height: 0;
      opacity: 0;
      margin-top: 0;
    }
    100% {
      margin-top: 20px;
      height: 40px;
      opacity: 1;
    }
  }
</style>
