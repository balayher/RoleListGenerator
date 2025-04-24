# RoleListGenerator

Generates a randomized custom role list for the ISFL ToS Server.

## Prerequisites

Before running this project locally, ensure you have the following installed:

- Go (version 1.24.0 or above)

## Installations

Download the files in this repository into a single folder. Enter 'go build' in the terminal to compile the program. Then enter 'go run RoleListGenerator' to run the program.

## How to Use

When starting the program, you will be prompted to input the desired number for each category. This value *must* be an integer, otherwise it will default to 0. Floating numbers will *not* work. The categories include:

- Town Investigative
- Town Protective
- Town Support
- Town Killing
- Random Town
- Mafia Killing
- Mafia Support
- Mafia Deception
- Random Mafia
- Coven Evil
- Vampires
- Neutral Killing
- Neutral Chaos
- Neutral Evil
- Neutral Benign
- Random Neutral
- Any

Depending upon your inputs, you may be asked some of these follow-up yes/no questions. Entering anything other than y, yes, n, or no will result in their default value. These situations are:

- If you have any Town Killing or Random Town, you will be asked if you want a guaranteed Jailor *(default: no)*
- If you have any Mafia Killing or Random Mafia, you will be asked if you want a guaranteed Godfather *(default: no)*
- If you have any Coven Evil, you will be asked if you want a guaranteed Coven Leader *(default: no)*
- If you have any Neutral Chaos, Random Neutral, or Any, you will be asked if you want Vampires in the random pool *(default: yes)*
- If you do not have any Mafia roles, you will be asked if you want Mafia in the Any pool *(default: yes)*
- If you do not have any Coven Evil, you will be asked if you want Coven in the Any pool *(default: yes)*
- You will be asked if you want to use the custom roles added by the ISFL community. Saying no will only give you standard Town of Salem roles *(default: yes)*

The last prompt will ask for any roles you want to ban. Each role must be separated by a space. Any role with multiple words must have a _ connecting the words (i.e. Coven_Leader). If you don't want any roles banned, you may leave this blank. Banned roles will overwrite guaranteed roles and Vampires if they were previously chosen.

There are also a few unique scenarios that may influence the randomization:

- If you add Mafia Support or Mafia Deception but no Mafia Killing or Random Mafia, a single Mafia Killing will be added to the role list to ensure that Mafia has either a Godfather or Mafioso.
- Similarly, if there no Mafia were initially added but at least one is rolled in an Any slot, if no Mafioso or Godfather is present in the final list, the first rolled Mafia role will be rerolled into either Godfather or Mafioso.
- If no Vampires were added initially or rolled during in a Neutral Chaos or Random Neutral slot, Vampire Hunter will be removed as an option from Town and Any rolls.
- If no Mafia or Coven were initially added, the Turncoat variant for that faction will be removed as an option from Neutral and Any rolls.
- If either Vampire, Mafia, or Coven is rolled in an Any slot after these removals, the corresponding Vampire Hunter or Turncoat (if custom roles are turned on) will be readded as an option for the remaining Any rolls (this will *not* apply retroactively to previous Town, Neutral, or Any rolls).
- If you add Coven Evil or allow Coven to appear in an Any slot, Witch will be removed as an option from Neutral and Any rolls.
- If no roles are available before starting generation of a particular category, those roles will be converted to the next inclusive level (i.e. Town Support will convert to Random Town, Random Town will convert to Any). If the available roles run out *during* generation of a category, the remaining roles for that category will be removed.

## Credits

Thanks to the members of the ISFL and the ISFL ToS Server for creating a great community.
Thanks to Broda for brainstorming ideas.
