# RoleListGenerator

Generates a randomized custom role list for the ISFL Town of Salem Discord server.

## Prerequisites

Before running this project locally, ensure you have the following installed:

- [Go](https://go.dev/doc/install) (version 1.24.0 or above)

## Installations

1. Ensure you have a valid version of [Go](https://go.dev/doc/install) installed. 
2. Download the files in this repository into a single folder. 
3. Navigate to that folder in your terminal and compile the program. 
``` bash
go build
```
4. The Role List Generator is now ready to run!
``` bash
go run RoleListGenerator
```

## How to Use

When starting the program, you will be prompted to input your desired number of roles for each category. This value *must* be an integer, otherwise it will default to 0. Floating numbers will *not* work. The categories include:

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

Depending upon your inputs, you may be asked some of these follow-up yes or no questions. Entering anything other than y, yes, n, or no will result in the default value listed below. These questions include:

- If you have any Town Killing or Random Town, you will be asked if you want a guaranteed Jailor *(default: no)*
- If you have any Mafia Killing or Random Mafia, you will be asked if you want a guaranteed Godfather *(default: no)*
- If you have any Coven Evil, you will be asked if you want a guaranteed Coven Leader *(default: no)*
- If you have any Neutral Chaos, Random Neutral, or Any, you will be asked if you want to allow Vampires in the random pools *(default: yes)*
- If you do not have any Mafia roles, you will be asked if you want Mafia in the Any pool *(default: yes)*
- If you do not have any Coven Evil, you will be asked if you want Coven in the Any pool *(default: yes)*
- You will be asked if you want to use the custom roles added by the ISFL community. Saying no will only give you standard Town of Salem roles *(default: yes)*

The last input prompt will ask for any roles you want to ban. Each role must be separated by a space. Any role with multiple words must have a _ connecting the words (i.e. 'Coven_Leader'). Turncoat will need it's corresponding faction specified (i.e. 'Turncoat(Mafia)' or 'Turncoat(Coven)'). If you don't want any roles banned, you may leave this blank. Banned roles will overwrite guaranteed roles and Vampires if they were previously chosen.

In addition, you will have two options for receiving your roles:

- You will be asked if you want the roles numbered to allow for easier assignment *(default: yes)*
- You will be asked if you want the rolelist to be written to roles.txt. If not, they will be printed directly to the terminal *(default: no)*

There are a few scenarios that may influence the randomization:

- If you add Mafia Support or Mafia Deception but not Mafia Killing or Random Mafia, one Mafia Support or Deception will be converted into Mafia Killing to ensure that Mafia has either a Godfather or Mafioso.
- Similarly, if no Mafia roles were added upfront but at least one is rolled in an Any slot, the first Mafia role will be rerolled into either Godfather or Mafioso if there isn't one present in the final role list.
- If no Vampires were added upfront or rolled in a Neutral Chaos or Random Neutral slot, Vampire Hunter will be removed as an option from Town and Any rolls.
- If no Mafia or Coven were added upfront, the Turncoat variant for that faction will be removed as an option from Neutral and Any rolls.
- If no Town were added upfront, Executioner will be removed as an option from Neutral and Any rolls.
- If an Executioner is rolled but non of the Town roles selected are eligible to have an Executioner, all Executioners will be changed to Jesters.
- If either Vampire, Mafia, Coven, or an Executioner viable Town is rolled in an Any slot after these removals, the corresponding Vampire Hunter, Turncoat (if custom roles are turned on), or Executioner role will be readded as an option for the remaining Any rolls if they are not banned (this will *not* apply retroactively to prior rolls).
- If you add Coven Evil or allow Coven to appear in an Any slot, Witch will be removed as an option from Neutral and Any rolls.
- If no roles are available before starting generation of a particular category, those roles will be converted to the next inclusive level (i.e. Town Support will convert to Random Town, Random Town will convert to Any, etc.). If the available roles run out *during* generation of a category, the remaining roles for that category will be removed.

## Credits

Thanks to the members of the ISFL and the ISFL ToS Server for creating a great community.
Thanks to Broda for brainstorming ideas.
