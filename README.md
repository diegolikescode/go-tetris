# DUMMY TEST

## HOW TO RUN:

just run ``make`` in the backend and ``yarn dev`` in the frontend and you will have the application up and running. The ``make`` is also using yarn, if you'd like to run ``make`` with npm just run ``make npm`` and that's about it. If you have ``nvm`` installed, run ``nvm use`` in the root folder where it has the ``.nvmrc`` file and you will get the NodeJS version I used for the test.

## BACKEND INTRODUCTION

I would like to explain my decision making for both front and backend

- I used sqlite with Sequelize as my ORM instead of reading from the CSV, this parsing was a nice bonus and I believe it is a real use case, so I included this and completed the requirements as they were stated;
- for the rest, I just think my Error Handling could be better, but I did it in kinda of a rush, so I'm sorry for any bad code you read!


## FRONTEND INTRODUCTION

- as for the frontend, let me be the first to tell you that the color scheme is called ``**DEBUG_MODE**``;
- as well as the backend, here the error handling isn't the best, but again, I was running with another side project for the end of my last semester in college (aka TCC);
- a bunch of ``any`` in my TS typing, which I'm not proud of and could have done better;
- the screen re-render more than I'd like, but it's not that bad;
- instead of using a CSS framework such as Material UI or Tailwind, I used ``styled-componens``, so I wrote raw CSS basically, that was a bad choise since Material UI or Tailwind helps a lot with responsiveness, and ``styled-components`` is more like "writing CSS in a TS/JS file" (I could be wrong, but I readed the docs), but at least I enjoy writing CSS, so the responsiveness is pretty bad, honestly;
