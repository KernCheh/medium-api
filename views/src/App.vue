<template>
  <div id="app" class="blogr">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="scss">
.blogr {
//   font-family: sans-serif;

//   a {
//     border: none;
//   }

//   .blogr-cont {
//     max-width: 960px;
//     width: 95%;
//     margin: 0 auto;
//   }

//   .posts-title {
//     h1 {
//       float: left;
//       font-size: 2.5em;
//       margin-bottom: 1em;
//     }

//     a {
//       float: right;
//       display: inline-block;
//       color: #CE0A0A;
//       border: 1px solid;
//       border-radius: 2em;
//       padding: .5em 1.4em;
//       margin-top: 1em;
//     }
//   }

//   .posts-nav {
//     border-top: 1px solid #ddd;
//     margin-bottom: 2em;

//     li {
//       display: inline-block;
//       line-height: 5;
//       text-transform: uppercase;
//       font-size: .8em;
//       letter-spacing: 2px;
//       font-weight: 700;
//       margin-right: 2em;
//       margin-top: -1px;
//       color: #222;
//     }

//     span {
//       display: block;
//       border-top: 1px solid #666;
//     }

//     a {
//       color: #999;
//     }
//   }

//   .blogr-posts {
//     .summary {
//       margin-bottom: 3em;

//       h3 {
//         a {
//           font-family: serif;
//           font-size: 2em;
//           color: #222;
//         }
//       }

//       .desc {
//         color: #aaa;

//         a {
//           color: #aaa;
//         }
//       }

//       .dot {
//         position: relative;
//         top: -.2em;
//         display: inline-block;
//         padding: 0 .3em;
//       }
//     }
//   }

  .blogr-form {
    font-family: serif;
    max-width: 50em;

    textarea {
      display: none;
    }

    .title {
      font-size: 2em;
      color: #222;
      font-weight: bold;
      margin-bottom: 1em;
    }

    .body {
      font-size: 1.3em;
    }
  }
}

// @-webkit-keyframes medium-editor-image-loading {
//   0% {
//     -webkit-transform: scale(0);
//             transform: scale(0); }
//   100% {
//     -webkit-transform: scale(1);
//             transform: scale(1); } }

// @keyframes medium-editor-image-loading {
//   0% {
//     -webkit-transform: scale(0);
//             transform: scale(0); }
//   100% {
//     -webkit-transform: scale(1);
//             transform: scale(1); } }

// @-webkit-keyframes medium-editor-pop-upwards {
//   0% {
//     opacity: 0;
//     -webkit-transform: matrix(0.97, 0, 0, 1, 0, 12);
//             transform: matrix(0.97, 0, 0, 1, 0, 12); }
//   20% {
//     opacity: .7;
//     -webkit-transform: matrix(0.99, 0, 0, 1, 0, 2);
//             transform: matrix(0.99, 0, 0, 1, 0, 2); }
//   40% {
//     opacity: 1;
//     -webkit-transform: matrix(1, 0, 0, 1, 0, -1);
//             transform: matrix(1, 0, 0, 1, 0, -1); }
//   100% {
//     -webkit-transform: matrix(1, 0, 0, 1, 0, 0);
//             transform: matrix(1, 0, 0, 1, 0, 0); } }

// @keyframes medium-editor-pop-upwards {
//   0% {
//     opacity: 0;
//     -webkit-transform: matrix(0.97, 0, 0, 1, 0, 12);
//             transform: matrix(0.97, 0, 0, 1, 0, 12); }
//   20% {
//     opacity: .7;
//     -webkit-transform: matrix(0.99, 0, 0, 1, 0, 2);
//             transform: matrix(0.99, 0, 0, 1, 0, 2); }
//   40% {
//     opacity: 1;
//     -webkit-transform: matrix(1, 0, 0, 1, 0, -1);
//             transform: matrix(1, 0, 0, 1, 0, -1); }
//   100% {
//     -webkit-transform: matrix(1, 0, 0, 1, 0, 0);
//             transform: matrix(1, 0, 0, 1, 0, 0); } }

// Medium editor
.medium-editor-anchor-preview {
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size: 16px;
  left: 0;
  line-height: 1.4;
  max-width: 280px;
  position: absolute;
  text-align: center;
  top: 0;
  word-break: break-all;
  word-wrap: break-word;
  visibility: hidden;
  z-index: 2000; }
  .medium-editor-anchor-preview a {
    color: #fff;
    display: inline-block;
    margin: 5px 5px 10px; }

.medium-editor-anchor-preview-active {
  visibility: visible; }

.medium-editor-dragover {
  background: #ddd; }

.medium-editor-image-loading {
  -webkit-animation: medium-editor-image-loading 1s infinite ease-in-out;
          animation: medium-editor-image-loading 1s infinite ease-in-out;
  background-color: #333;
  border-radius: 100%;
  display: inline-block;
  height: 40px;
  width: 40px; }

.medium-editor-placeholder {
  position: relative; }
  .medium-editor-placeholder:after {
    content: attr(data-placeholder) !important;
    font-style: italic;
    position: absolute;
    left: 0;
    top: 0;
    white-space: pre;
    padding: inherit;
    margin: inherit; }

.medium-editor-placeholder-relative {
  position: relative; }
  .medium-editor-placeholder-relative:after {
    content: attr(data-placeholder) !important;
    font-style: italic;
    position: relative;
    white-space: pre;
    padding: inherit;
    margin: inherit; }

.medium-toolbar-arrow-under:after, .medium-toolbar-arrow-over:before {
  border-style: solid;
  content: '';
  display: block;
  height: 0;
  left: 50%;
  margin-left: -8px;
  position: absolute;
  width: 0; }

.medium-toolbar-arrow-under:after {
  border-width: 8px 8px 0 8px; }

.medium-toolbar-arrow-over:before {
  border-width: 0 8px 8px 8px;
  top: -8px; }

.medium-editor-toolbar {
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size: 16px;
  left: 0;
  position: absolute;
  top: 0;
  visibility: hidden;
  z-index: 2000; }
  .medium-editor-toolbar ul {
    margin: 0;
    padding: 0; }
  .medium-editor-toolbar li {
    float: left;
    list-style: none;
    margin: 0;
    padding: 0; }
    .medium-editor-toolbar li button {
      box-sizing: border-box;
      cursor: pointer;
      display: block;
      font-size: 14px;
      line-height: 1.33;
      margin: 0;
      padding: 15px;
      text-decoration: none; }
      .medium-editor-toolbar li button:focus {
        outline: none; }
    .medium-editor-toolbar li .medium-editor-action-underline {
      text-decoration: underline; }
    .medium-editor-toolbar li .medium-editor-action-pre {
      font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
      font-size: 12px;
      font-weight: 100;
      padding: 15px 0; }

.medium-editor-toolbar-active {
  visibility: visible; }

.medium-editor-sticky-toolbar {
  position: fixed;
  top: 1px; }

.medium-editor-relative-toolbar {
  position: relative; }

.medium-editor-toolbar-active.medium-editor-stalker-toolbar {
  -webkit-animation: medium-editor-pop-upwards 160ms forwards linear;
          animation: medium-editor-pop-upwards 160ms forwards linear; }

.medium-editor-action-bold {
  font-weight: bolder; }

.medium-editor-action-italic {
  font-style: italic; }

.medium-editor-toolbar-form {
  display: none; }
  .medium-editor-toolbar-form input,
  .medium-editor-toolbar-form a {
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif; }
  .medium-editor-toolbar-form .medium-editor-toolbar-form-row {
    line-height: 14px;
    margin-left: 5px;
    padding-bottom: 5px; }
  .medium-editor-toolbar-form .medium-editor-toolbar-input,
  .medium-editor-toolbar-form label {
    border: none;
    box-sizing: border-box;
    font-size: 14px;
    margin: 0;
    padding: 6px;
    width: 316px;
    display: inline-block; }
    .medium-editor-toolbar-form .medium-editor-toolbar-input:focus,
    .medium-editor-toolbar-form label:focus {
      -webkit-appearance: none;
         -moz-appearance: none;
              appearance: none;
      border: none;
      box-shadow: none;
      outline: 0; }
  .medium-editor-toolbar-form a {
    display: inline-block;
    font-size: 24px;
    font-weight: bolder;
    margin: 0 10px;
    text-decoration: none; }

.medium-editor-toolbar-form-active {
  display: block; }

.medium-editor-toolbar-actions:after {
  clear: both;
  content: "";
  display: table; }

.medium-editor-element {
  word-wrap: break-word;
  min-height: 30px; }
  .medium-editor-element img {
    max-width: 100%; }
  .medium-editor-element sub {
    vertical-align: sub; }
  .medium-editor-element sup {
    vertical-align: super; }

.medium-editor-hidden {
  display: none; }
</style>
