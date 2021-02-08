<script>
</script>

<main>
    <h1>Megami 女神</h1>
    <p>The cross-platform MyAnimeList client.</p>
    <p>Written in Electron.js + Svelte</p>

    <div class="login-form">
        <form>
            <label for="mal-username">Username</label><input id="mal-username"/>
            <label for="mal-password">Password</label><input id="mal-password" type="password"/>
            <button type="submit">Log In</button>
        </form>
    </div>
</main>

<style type="text/scss">
  @import './colours.scss';

  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  div.login-form {
    background-color: $button;
    display: inline-block;
    padding: 1rem;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>