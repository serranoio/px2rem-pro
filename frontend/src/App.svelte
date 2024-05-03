<script lang="ts">
  let output: {contents: string } = {contents: "output"};

  let myInput: HTMLInputElement;
  const onClick = async () => {
      const textarea = document.querySelector("textarea")

      const hosted = "https://px2rem-pro-divine-sky-9836.fly.dev/convert"
      // const local = "http://localhost:6883/convert"
        
        try {
          const data = await fetch(hosted, {
            method: "POST",
            body: JSON.stringify({
              conversionFactor: myInput?.value,
              input: textarea?.value, 
            })
          })
          
          output = await data.json()
        
      
    } catch(err) {

      output.contents = "Please insert a value in conversion factor"
    }


  }
</script>

<main>
  <div class="github-logo">
    <a target="_blank" href="https://github.com/serranoio/px2rem-pro">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256"><rect width="256" height="256" fill="none"/><path d="M119.83,56A52,52,0,0,0,76,32a51.92,51.92,0,0,0-3.49,44.7A49.28,49.28,0,0,0,64,104v8a48,48,0,0,0,48,48h48a48,48,0,0,0,48-48v-8a49.28,49.28,0,0,0-8.51-27.3A51.92,51.92,0,0,0,196,32a52,52,0,0,0-43.83,24Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="16"/><path d="M104,232V192a32,32,0,0,1,32-32h0a32,32,0,0,1,32,32v40" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="16"/><path d="M104,208H72a32,32,0,0,1-32-32A32,32,0,0,0,8,144" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="16"/></svg>
    </a>
  </div>
  <h1>Pixel 2 Rem Pro</h1>
  <div class="center">
    <label>Conversion Factor:</label>
    <input value="0" placeholder=".5" bind:this={myInput} type="number"/>
  </div>

  <figure>
    <div class="input">
      <textarea placeholder="Input: " />
    </div>
    <div class="output">
      {#each output.contents.split("\n") as line}
        <p>
          {line}
        </p>
      {/each}
    </div>
  </figure>

<div class="center">
  <button on:click={onClick} class="convert">Convert</button>
</div>

</main>

<style>
  main {
    position: relative;
  }
  .github-logo {
    position: absolute;
    right: 1%;
    top: 1%;
  }
  .github-logo svg {
    width: 3rem;
    height: 3rem;
  }
  .github-logo svg path {
    stroke: var(--gray-70l);
  }

  .convert {
    padding: 1.2rem 2.4rem;    
    margin-top: 2rem;
    border-radius: 2.4rem;
    border: none;
    color: var(--gray-90l);
    font-size: 2rem;
    background-image: linear-gradient(to left, var(--primary-1), var(--primary-2), var(--primary-3), var(--primary-4));
    transition: all .2s;
    cursor: pointer;
    box-shadow: 0 2px 2px 4px rgba(0,0,0,.2);
  }
  
  .convert:hover {
    transform: translateY(2px);
    box-shadow: 0 0 2px 2px rgba(0,0,0,.2);
  }
  
  label {
    color: var(--gray-80l);
    font-size: 2rem;
  }
  .center {
    display: flex;
    justify-content: center;
    align-items: center;
    padding-bottom: 2rem;
    gap: 1.2rem;
    font-size: 1.2rem;
  }
  
  input {
    height: 2.4rem;
    border: none;
    background-color: var(--gray-90l);
    padding: .6rem;
    border-radius: 2rem;
    font-size: 2rem;
    width: 10%;
    color: var(--primary-3);
  }
  
  input::placeholder {
    color: var(--gray-70l);
  }
  
  h1 {
    color: var(--gray-90l);
    font-size: 5.2rem;
    text-align: center;
    padding: 2.4rem;
  }
  
  h2 {
    font-size: 1.6rem;
    text-align: left;
  }
  figure {
    display: flex;
    gap: 2.4rem;
    height: 60vh;
    padding: 2.4rem;
    max-width: 90vw;
    margin: 0 auto;
    border-radius: 2.4rem;
    background-image: linear-gradient(to left, var(--primary-1), var(--primary-2), var(--primary-3), var(--primary-4));
    box-shadow: 0 2px 2px 4px rgba(0,0,0,.2);
  }
  
.input, .output {
  padding: 2rem;
  border-radius: 2.4rem;
  font-size: 2.4rem;
  width: 100%;
  background-color: var(--gray-90l);
  color: var(--primary-2);
}

.output {
  overflow-y: scroll;
  font-size: 1.6rem;
}

textarea, .output {
  font-size: 1.6rem;
}
textarea {
  color: var(--primary-1);
  width: 100%;
  height: 100%;
  border: none;
  resize: none;
  background-color: transparent;
}

textarea::placeholder {
  font-size: 1.6rem;
  position: absolute;
  top: 1%;
  left: 1%;
  color: var(--gray-70l);
}




</style>
