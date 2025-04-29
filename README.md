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

When starting the program, you will be asked if you want to load your role counts from a .json file. The role counts are the number of roles to be selected from each role category. If yes, you will be prompted to enter the name of the .json file to use. The structure of this .json file will be detailed below. If you don't want to load from a file (or if there's an error loading from the file), you will be asked for the desired number of roles for each category. This value *must* be an integer, otherwise it will default to 0. Floating numbers will *not* work. You will also have the option to save your role counts to counts.json for future use (note that this will overwrite your current counts.json file). The categories include:

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

Next, you will be asked if you want to load your role options from a .json file. If yes, you will be prompted to enter the name of the .json file to use. The structure of this .json file will be detailed below. If you don't want to load from a file (or if there's an error when loading from the file), you will be prompted for each option. Entering anything other than y, yes, n, or no will result in the default value listed below. You will also have the option to save your options to options.json for future use (note that this will overwrite your current options.json file). These options include:

- Do you want a guaranteed Jailor? *(default: no)*
- Do you want a guaranteed Godfather? *(default: no)*
- Do you want a guaranteed Coven Leader? *(default: no)*
- Do you want Vampires in the random pools? *(default: yes)*
- Do you want Mafia in the Any pool? *(default: yes)*
- Do you want Coven in the Any pool? *(default: yes)*
- Do you want to use the custom roles added by the ISFL community? If not, you will only get standard Town of Salem roles. *(default: yes)*
- Do you want to number the roles in the final role list? *(default: yes)*
- Do you want the rolelist to be written to roles.txt? If not, they will be printed directly to the terminal (note that this will overwrite your current roles.txt file). *(default: no)*

You will also be asked if you want to ban any role. Each role must be separated by a space. Any role with multiple words must have a _ connecting the words (i.e. 'Coven_Leader'). Turncoat will need it's corresponding faction specified (i.e. 'Turncoat(Mafia)' or 'Turncoat(Coven)'). If you don't want any roles banned, you may leave this blank. Banned roles may override any guaranteed option chosen.

There are a few scenarios that may influence the randomization:

- If you add Mafia Support or Mafia Deception but not Mafia Killing or Random Mafia, one Mafia Support or Deception will be converted into Mafia Killing to ensure that Mafia has either a Godfather or Mafioso.
- Similarly, if no Mafia roles were added upfront but at least one is rolled in an Any slot, the first Mafia role will be rerolled into either Godfather or Mafioso if there isn't one present in the final role list.
- If no Vampires were added upfront or rolled in a Neutral Chaos or Random Neutral slot, Vampire Hunter will be removed as an option from Town and Any rolls.
- If no Mafia or Coven were added upfront, the Turncoat variant for that faction will be removed as an option from Neutral and Any rolls.
- If no Town were added upfront, Executioner will be removed as an option from Neutral and Any rolls.
- If an Executioner is rolled but no viable Executioner targets are rolled, all Executioners will be converted to Jesters (even if Jester was banned).
- If a Guardian Angel is rolled but no viable Guardian Angel targets are rolled, all Guardian Angels will be converted to Survivors (even if Survivor was banned).
- If either Vampire, Mafia, Coven, or an Executioner viable Town is rolled in an Any slot after these removals, the corresponding Vampire Hunter, Turncoat (if custom roles are turned on), or Executioner role will be readded as an option for the remaining Any rolls if they are not banned (this will *not* apply retroactively to prior rolls).
- If you add Coven Evil or allow Coven to appear in an Any slot, Witch will be removed as an option from Neutral and Any rolls.
- If no roles are available before starting generation of a particular category, those roles will be converted to the next inclusive level (i.e. Town Support will convert to Random Town, Random Town will convert to Any, etc.). If the available roles run out during generation of a category, the remaining roles will be converted to the next inclusive level. If there are no more Any roles available, the remaining slots are removed (this should only occur if you ban all non-unique roles from the rolelist).

## Json File Formats

### Role Count

To load in a role count, your .json file must be in the following format:

```
{
    "ti": 3,
    "tp": 2,
    "ts": 1,
    "tk": 1,
    "rt": 3,
    "mk": 1,
    "ms": 1,
    "md": 1,
    "rm": 1,
    "ce": 4,
    "nk": 1,
    "nc": 0,
    "ne": 0,
    "nb": 0,
    "rn": 2,
    "a": 1,
    "vamp": 0
}
```

All variables must be integers or the load will fail. Any missing variable will be set to 0.

These variables correspond to the following role categories:

```
ti: Town Investigative
tp: Town Protective
ts: Town Support
tk: Town Killing
rt: Random Town
mk: Mafia Killing
ms: Mafia Support
md: Mafia Deception
rm: Random Mafia
ce: Coven Evil
nk: Neutral Killing
nc: Neutral Chaos
ne: Neutral Evil
nb: Neutral Benign
rn: Random Neutral
a: Any
vamp: Vampires
```

### Role Options

To load in role options, your .json file must be in the following format:

```
{
	"jailor": false,
	"gf": false,
	"cl": false,
	"anyMaf": true,
	"anyCov": true,
	"anyVamp": true,
	"custom": true,
	"numbered": true,
	"fileWrite": false
}
```

All variables must be booleans or the load will fail. Any missing variable will be set to false.

These variables correspond to the following options:

```
jailor: Do you want a guaranteed Jailor?
gf: Do you want a guaranteed Godfather?
cl: Do you want a guaranteed Coven Leader?
anyMaf: Do you want Mafia in the Any pool?
anyCov: Do you want Coven in the Any pool?
anyVamp: Do you want Vampires in any random pools?
custom: Do you want custom roles in the pool?
numbered: Do you want your role list output to be numbered?
fileWrite: Do you want the results written to roles.txt?
```

## Credits

Thanks to the members of the ISFL and the ISFL ToS Server for creating a great community.
Thanks to Broda for brainstorming ideas.
