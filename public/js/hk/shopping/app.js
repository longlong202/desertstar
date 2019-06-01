/* default dom id (particles-js) */
//particlesJS();

/* config dom id */
//particlesJS('dom-id');

/* config dom id (optional) + config particles params */
particlesJS('particles-js', {
  particles: {
      number: {
          value: 120,
          density: {
              enable: true,
              value_area: 800
          }
      },
      color: '#951de7',
      shape: 'circle',
      opacity: 1,
      size: 4,
      size_random: true,
      line_linked: {
          enable_auto: true,
          distance: 140,
          color: '#951de7',
          opacity: 1,
          width: 1,
          condensed_mode: {
              enable: true,
              rotateX: 1000,
              rotateY: 1000
          }
      },
      move: {
          enable: true,
          speed: 6,
          direction: "none",
          random: false,
          straight: false,
          out_mode: "out",
          attract: {
              enable: false,
              rotateX: 600,
              rotateY: 1200
          }
      }
  },
  interactivity: {
      detect_on: 'canvas',
      events: {
          onclick: {
              enable: true,
              mode: "push"
          },
          resize: true
      },
      modes: {
          push: {
              particles_nb: 4
          }
      }
  },
  /* Retina Display Support */
  retina_detect: true
});