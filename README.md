# Gestione studenti

La repo e' stata implementata in momenti molto diversi e nel corso del tempo la repo e' stata usata come cavia per testare diverse tecnologie, 
quindi si e' pregati di non commentare troppo essendo che ci sono almeno 3 logiche di gestione diverse **perche' per testare ho usato**
**i paradigmi soliti delle teconogie usate**.

# ULTIMATE Stack
Backend:
- Go (with the new goat std library net/http e go > 1.22.x)
- [templ](https://templ.guide/syntax-and-usage/context/) (for my juicy html, ehm... sorry **HTMX**)

Db:
- Postgress (the best relational db <3) {mogno is for mogols that are sfaticated}
- Redis (after all why whe need percistancy???)

Frontend:
- htmx (The real goat backend-frontend library. Why two dom when we can have no dom?)

Development tools:
- the little whale [docker] (but i prefer lxc for containerization) {rich wold approve, as he said "gnu is not unix"}
- pyhon (import requests)
- air (rebild automatically when a file is changed)

## Come avviare l'applicazione
- Creare un progetto nella console di [google cloud](https://console.cloud.google.com).
- Andare nella sezione credenziali e creare delle credenziali per la voce `ID client OAuth 2.0`.
- Nella voce `Origini JavaScript autorizzate` inserire l'url del sito (possibilmente con https, ma non necessario).
    [esempio: *http://localhost:8080*]
- Nella voce `URI di reindirizzamento autorizzati` inserire l'ur del sito (come quello sopra) + `/auth/callback`.
    [esempio: *http://localhost:8080/auth/callback*]
- Mettere nel file `.env` alla voce `REDIRECT_URL` il link inserito nella voce `URI di reindirizzamento autorizzati` della console di google.
- Per avviare l'applicazione avviare docker e fare il comando `docker compose up`.