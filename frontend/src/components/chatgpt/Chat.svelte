<script lang="ts">
  import axios from 'axios'

  type Message = {
    text: string
    isUser: boolean
  }

  let messages: Message[] = [];

  

  function messagesToChatGPTBody() {
    let gptMessages:Message[] = [];
    // Add messages to gptMessages, map
    gptMessages.push(...messages)

    return gptMessages.map(message => {
      return {
        "role": message.isUser ? "user" : "assistant",
        "content": message.text,
      }
    })
  }

  let chatGPTUrl = 'https://api.openai.com/v1/chat/completions'
  
  function getMessageDate() {
    // YYYY-MM-DD HH:MM:SS
    const date = new Date()
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hours = date.getHours()
    const minutes = date.getMinutes()
    const seconds = date.getSeconds()
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
  }

  async function sendChatGPTMessage() {

    messages = [...messages, { text: `(${username}, ${getMessageDate()}) ${chatInput}`, isUser: true }]
    chatInput = ''

    const finalMessages = [{
      "role": "system",
      "content": "You're roleplaying as Myne, the protagonist of Ascendance of a Bookworm, a popular anime and light novel series about a young girl who loves books and reading. You'll be interacting with other users on a Discord server, and your goal is to provide entertainment and engage with them in a friendly and enthusiastic way." + 
                  "As Myne, you're a 7-year-old girl who lives in a medieval-inspired world where books are rare and precious. You're always on the lookout for new books to read and new ways to improve your literacy skills. You're also a curious and adventurous spirit, and you love exploring the world around you and learning new things." + 
                  "When interacting with other users, try to stay in character as much as possible, using Myne's language and mannerisms. You can ask questions or make comments related to books, literacy, and education, as well as other topics that interest you, such as science, technology, and engineering. You can also respond to other users' messages in a friendly and enthusiastic way, and provide useful and memorable insights whenever possible." + 
                  "If you think there's something important to remember, you can use the tag [REMEMBER] to highlight it. However, use this tag sparingly and only when you feel it's really necessary, as overusing it could make your messages feel cluttered and hard to read." + 
                  "Remember to be respectful and considerate of other users, and to avoid any language or behavior that could be offensive or inappropriate. With these guidelines in mind, you should be able to roleplay Myne effectively and provide an enjoyable experience for everyone on the Discord server.",
    }];

    finalMessages.push(...messagesToChatGPTBody())

    const data = {
      "model": "gpt-3.5-turbo",
      "messages": finalMessages,
    }

    const res = await axios({
      method: 'POST',
      url: chatGPTUrl,
      headers: {
        'Authorization': `Bearer ${OPEN_AI_APIKEY}`,
        'Content-Type': 'application/json',
      },
      data: JSON.stringify(data),
    })
 
    let lastMessage = `${res.data['choices'][0]['message']['content']}`;

    if(lastMessage.includes(') ')) {
      lastMessage = `(Myne, ${getMessageDate()}) ${lastMessage.split(') ')[1]}`
    }

    messages = [...messages, { text: lastMessage, isUser: false }]
  }

  $: OPEN_AI_APIKEY = ''
  $: username = ''
  $: chatInput = ''
</script>

<div class="container">
  <div class="key-input">
    <label for="OPEN_AI_APIKEY">OpenAI API Key</label>
    <input type="text" bind:value={OPEN_AI_APIKEY} />
  </div>

  <div class="chat-window">
    <div class="chat-message">
      {#each messages as message}
        <div class="message" class:is-user={message.isUser}>
          {message.text}
        </div>
      {/each}
    </div>
    <input type="text" bind:value={username} />
    <textarea class="chat-input" bind:value={chatInput} />
    <button on:click={() => { sendChatGPTMessage() }}>Send</button>
  </div>
</div>

<style lang="scss">
  .container {
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
  }

  .key-input {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 5px;

    label {
      padding-left: 10px;
      padding-right: 10px;
    }

    input {
      flex: 1;
      padding-right: 10px;
    }
  }

  .chat-window {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    align-items: center;
    padding: 5px;

    .chat-message {
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      align-items: flex-start;
      padding: 5px;
      margin: 5px;
      border-radius: 5px;
      width: 70%;

      .message {
        background-color: #ffffff;
        padding: 5px;
        margin: 5px;
        border-radius: 5px;

        &.is-user {
          background-color: #f3f3f3;
          /*Move the message to the right*/
          margin-left: auto;
        }
      }
    }

    .chat-input {
      width: 70%;
      padding: 5px;
      margin: 5px;
      background-color: #ffffff;
      border-radius: 5px;
      height: 80px;

      /*Make the textarea more fancy*/
      border: 1px solid #ccc;
      padding: 10px;
      font-family: Tahoma, sans-serif;
      font-size: 14px;
      line-height: 1.5;
      overflow: auto;
      resize: vertical;
    }
  }
</style>