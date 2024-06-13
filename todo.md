~~- fare tutto perche' funzioni anche il collegamentos Frequenta~~
~~- fare che si crea il db ogni volta che si connette (se non esiste)~~

~~- aggiungiere possibilita' di update:~~
    ~~= Student~~
    ~~- Class~~
~~- aggiugere possibilita' di delete:~~
    ~~- student~~
    ~~- class~~
    ~~- major~~
    ~~- studentclass~~

~~- far si che tutte le entita' si possano salvare nel db~~
    ~~- student~~
    ~~- class~~
    ~~- major~~
    ~~- studentclass~~

- ~~fare endopoint per aggiungere item~~
    ~~- student ~~
    ~~- class ~~
    ~~ - major ~~
    ~~- studentclass~~
    - (Non penso che servi) user

- ~~fare endopoint per eliminare item~~
    ~~- student~~
    ~~- class~~
    ~~- major ~~
    ~~- studentclass~~
    - (non penso che servi) user

- ~~fare endopoint per modificare item~~
    ~~- student ~~
    ~~- class~~
    ~~- major ~~
    - user // mettere che possa farlo solo chi ha il ruolo admin

- ~~verificare che il token sia aggiornato (IsAdmin - IsEditor)~~
- ~~implementazione gestione sessioni~~
    ~~- creazione account ad un autenticazione con gogle~~
        ~~- una volta account creato aggiongiere modo per autorizzare~~
        ~~- reindirizzamento a pagina che dice di aspettare di essere autorizzati~~
    ~~- implementare bene pictures (in user)~~
    ~~- cambiare la il link di google da /auth/login/google => /login~~

PRIMA DI DELIVER METTI TUTTO IN UN FILE CONF/ENV


~~- integrare login google~~
~~- aggiungere users~~
~~- implementare un sistema per la gestione delle sessioni~~
- aggiustare quando si fa edit senza aver fatto create [non puo' esistere questa cosa]
~~- mettere funzioni sotto routingz~~
- negli errori dei services mettere la parola "fields"
- migrare services to handlers (metodi sensati)



Cose da contorllare:
- ~~una volta implementato il link tra studente e classe allora ~~
    ~~andare a testare la tabella di pagina "http://localhost:8080/students/1"~~
- ~~Poter modificare la classe di uno studente~~ ~~(ed anche vederla grazie)~~
- ~~quando si migra uno studente lo si vede 2 volte nella lista degli studenti della classe~~
    
- Aggiungiere un tasto per fare le migrazioni in automatico
- Controllare che quando aggiungo uno studente faccia vedere solo le classi attive

- aggiungiere un modo per eliminare le classi
- aggiuingiere un allert prima di eliminare le cose
- aggiungere un tasto per eliminare lo storico dello studente in '/students/:id'

Optionals 
- Quando si modifica la classe dello studente allora bisogna aggiungeire il dato
    anche alla funzione sotto








Chicche finali:
    - mettere logo della scuola al posto dell'emoji




Cose brutte che non ho voglia di aggiustare:
    - nelle pag html non uso i template per renderizzare i componenti quindi quando devo modificare
        un componente devo anche andare a modificarlo nel suo file separato. Non ho voglia di spendere
        troppo tempo a capire come farlo con i template. Se si vuole fixare questo problema lo si puo'
        fare abbastanza semplicemente usando https://templ.guide ma ormai ho iniziato ad usare i template
        e continuero ad usarli. 
