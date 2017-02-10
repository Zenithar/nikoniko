<style>
.mood-awesome {
  background-color: #1abc9c
}

.mood-good {
  background-color: #5c91df
}

.mood-moderate {
  background-color: #f1c40f
}

.mood-bad {
  background-color: #ff7a23
}

.mood-angry {
  background-color: #F72E1A
}
</style>

<section class="hero is-primary is-bold">
  <div class="hero-body">
    <div class="container">
      <p class="title">
        Welcome !
      </p>
      <p class="subtitle">
        Every days has his mood.
      </p>
    </div>
  </div>
</section>

<section class="section">
  <div class="container">
    <form class="form">
      {{ .xsrf_data }}
      <input type="hidden" name="email" value="{{ .email }}">
      <input type="hidden" name="expires" value="{{ .expires }}">
      <input type="hidden" name="signature" value="{{ .signature }}">
      <label class="label">Message</label>
      <p class="control">
        <textarea class="textarea" placeholder="Optionally insert your message here ..."></textarea>
      </p>
      <label class="label">Choose your mood</label>
      <input type="hidden" name="mood" value="">
    </form>
  </div>
  <div class="container">
    <div class="columns">

      <div class="column">
        <p class="controél">
          <button type="button" class="button is-fullwidth is-large mood-angry">
            <span class="icon is-large">
              <i class="fa fa-thumbs-o-down"></i>
            </span>
          </button>
        </p>
      </div>

      <div class="column">
        <p class="control">
          <button type="button" class="button is-fullwidth is-large mood-bad">
            <span class="icon is-large">
              <i class="fa fa-frown-o"></i>
            </span>
          </button>
        </p>
      </div>

      <div class="column">
        <p class="control">
          <button type="button" class="button is-fullwidth is-large mood-moderate">
          <span class="icon is-large">
            <i class="fa fa-meh-o"></i>
          </span>
        </button>
        </p>
      </div>

      <div class="column">
        <p class="control">
          <button type="button" class="button is-fullwidth is-large mood-good">
          <span class="icon is-large">
            <i class="fa fa-smile-o"></i>
          </span>
        </button>
        </p>
      </div>

      <div class="column">
        <p class="control">
          <button type="button" class="button is-fullwidth is-large mood-awesome">
            <span class="icon is-large">
              <i class="fa fa-thumbs-o-up"></i>
            </span>
          </button>
        </p>
      </div>

    </div>
  </div>
</section>
