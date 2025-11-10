# VALORANT STATS OVERLAY

> Lekki i prosty overlay dla Valoranta, który pokazuje Twój aktualny rank, punkty (RR) oraz wynik ostatniej gry w czasie rzeczywistym.

![Overlay Screenshot](https://github.com/user-attachments/assets/837be80d-96c0-4158-b771-720749b6ab6e)

## Co posiada taki overlay?

- Pokazuje **aktualną rangę** i punkty rankingowe (RR).
- Wyświetla **wynik ostatniej gry** (czy wygrałeś, ile MMR zdobyłeś/straciłeś).
- **Łatwy w instalacji** i lekki dla systemu – nie spowalnia gry.
- Obsługa w **czasie rzeczywistym** – statystyki zawsze aktualne.

## Instalacja

1. Pobierz [najnowszą wersję](https://github.com/Kartosowski/valorant-stats-overlay/releases/download/1.0.0/vs-overlay-1.0.zip) i rozpakuj do folderu.  
2. Otwórz plik `config.json` w edytorze tekstu i wpisz swój **username, tag oraz własny klucz**:

```json
{
    "apikey": "tutaj-twoj-klucz",
    "username": "tutaj-twoj-username",
    "tag": "tutaj-twoj-tag",
    "port": 2025
}
```

> ⚠️ Aby zdobyć klucz, dołącz do [HenrikDev Discord](https://discord.gg/JDpQcB7nzw), wygeneruj klucz i wklej go w `apikey`.

3. Uruchom plik `vs-overlay.exe` w folderze. Serwer lokalny wystartuje na porcie ustawionym w `config.json` (domyślnie 2025).  
4. W terminalu znajdziesz link, który wpisujesz w **komponencie przeglądarki OBS**, aby overlay działał na streamie.

![OBS Example 1](https://github.com/user-attachments/assets/9345f884-9dde-42a4-a2aa-c8a6ead7053a)
![OBS Example 2](https://github.com/user-attachments/assets/66a900ba-589b-462c-87b6-468f1316d986)
