<template>
  <v-container class="container-width">
  <section class="notFound">
    <div class="img">
      <img src="@/assets/backToTheHomepage.webp" alt="Back to the Homepage" />
      <img src="@/assets/Delorean.webp" alt="Delorean, Doc, and Marty McFly" />
    </div>
    <div class="text">
      <h1>404</h1>
      <h2>PAGE NOT FOUND</h2>
      <h3>BACK TO HOME?</h3>
      <a href="/" class="yes">YES</a>
      <a href="https://www.youtube.com/watch?v=G3AfIvJBcGo">NO</a>
    </div>
  </section>
</v-container>
  <container class="footer">
    <div class="clean">
      404 Page Provided by: Copyright (c) 2023 by
      <a class="clean" href=https://codepen.io/honeybadger2788/pen/oNzKzvy>JSON.girlify</a>
    </div>
  </container>
</template>

<style scoped>
/* Copyright (c) 2023 by JSON.girlify (https://codepen.io/honeybadger2788/pen/oNzKzvy) */
@import url('https://fonts.googleapis.com/css2?family=Press+Start+2P&display=swap');

* {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  font-family: 'Press Start 2P';
  color: #FFFFFF;
  text-align: center;
}

.clean {
  /* height: 100vh; */
  font-family: "Roboto", sans-serif;
  line-height: 1.5;
  font-size:0.7rem;
  text-rendering: optimizeLegibility;
}

.footer {
  position: fixed;
  bottom: 0;
}

.container-width {
  max-width: 1570px;
  margin: 0 auto;
  /* height: 80vh; */
}

section.notFound {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  margin: 0 5%;
  height: 90vh;
}

section.notFound h1 {
  color: red;
  font-size: 100px;
}

section.notFound h2 {
  font-size: 50px;
}

section.notFound h1,
h2,
h3 {
  margin-bottom: 40px;
}

div.text {
  height: 50vh;
}

div.text a {
  text-decoration: none;
  margin-right: 20px;
}

div.text a:hover {
  color: red;
  text-decoration: underline;
}

@media only screen and (max-width: 1570px) {
  section.notFound {
    flex-direction: column;
  }

.container-width {
  max-width: 500px;
  margin: 0 auto;
}

  section.notFound div.img img {
    width: 25vw;
    height: auto;
  }

  section.notFound h1 {
    font-size: 25px;
    margin: 0em;
    margin-top:2em;
  }

  section.notFound h2, h3 {
    font-size: 12.5px;
    margin: 0.5em;
  }


  div.text a:active {
    color: red;
    text-decoration: underline;
  }
}

@media only screen and (max-height: 1000px) {
  section.notFound {
    flex-direction: column;
    height: 80vh;
  }

  .container-width {
    margin: 0 auto;
    height: 50vh;
  }

  section.notFound div.img img {
    width: 25vw;
    height: auto;
  }

  section.notFound h1 {
    font-size: 25px;
    margin: 0em;
    margin-top:2em;
  }

  section.notFound h2, h3 {
    font-size: 12.5px;
    margin: 0.5em;
  }

  div.text a:active {
    color: red;
    text-decoration: underline;
  }
}
</style>
