// ** FEUTREN CLI ** //

// -------------------------------------------------------------
// Initialisation des constantes
const readline = require('readline');
const figlet = require("figlet");
const { spawn } = require('child_process');
const { exec } = require('child_process');
const command = spawn('bash');

// -------------------------------------------------------------
// Donner du style ANSI à l'affichage 
const rouge = '\x1b[31m%s\x1b[0m';
const cyan = '\x1b[36m%s\x1b[0m';
const bleuGras = '\u001b[1;34m';
const gras = '\u001b[1m';
const resetColor = '\u001b[0m';

// -------------------------------------------------------------
// Afficher le prompt

command.stdout.on('data', (data) => {
    console.log(rouge, `stdout: ${data}`);
});

command.stderr.on('data', (data) => {
  console.error(rouge, `stderr: ${data}`);
  rl.prompt();
});

command.on('close', (code) => {
  console.log(`child process exited with code ${code}`);
});

// -------------------------------------------------------------
// Afficher le texte stylisé

figlet('FEUTREN    CLI', function (err, data) {
    console.log(data);
    console.log("----------------------------------------");
    console.log(cyan, "'help' pour plus d'informations\n'exit' pour quitter le bash\n", resetColor);
    rl.prompt(); }) 
       
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

// Afficher ">>>"  Avant chaque ligne de commande
const prompt = '>>> ';

// -------------------------------------------------------------
// ** MAIN ** //

rl.on('line', (input) => {

  rl.setPrompt(prompt);
  rl.prompt();

  // EXIT
  if (input === 'exit') 
  {
    command.stdin.end();
    rl.close();
  } 

  // CLEAR
  else if (input === 'clear' || input === 'cls')
  {
    console.clear();
  }

  // HELP
  else if (input === 'help')
  {
    const { Command } = require('commander');
    const program = new Command();
    program
        .name('Feutren-cli')
        .description('Simple CLI en javascript')
        .version('0.1.0');
    program.command('clear')
        .description('Efface les lignes du terminal')
    program.command('ls')
        .description('Afficher les fichiers et dossiers')
    program.command('cd')
        .description('Changer de dossier')
    program.command('touch')
        .description('Créer un fichier')
    program.command('mkdir')
        .description('Créer un dossier')
    program.command('rm')
        .description('Supprimer un fichier ou un dossier')
    program.parse(process.argv);
  } 

  // LS
  else if (input === 'ls') {
    exec('ls -F', (err, stdout, stderr) => {
      if (err) {
        console.error(err);
        return;
      }
      if (stderr) {
        console.error(stderr);
      }
  
      const lignes = stdout.trim().split('\n');
      for (let ligne of lignes) {
        if (ligne.endsWith('/')) {
          console.log(bleuGras + ligne + resetColor);
        } else {
          console.log(gras + ligne + resetColor);
        }
      }
    });
  }

  // CD
  else if (input.startsWith('cd '))
  {
    const directory = input.slice(3);
    command.stdin.write(`cd ${directory}\n`);
  }

  // TOUCH
  else if (input.startsWith('touch ')) {
    const filename = input.slice(6);
    exec(`touch ${filename}`, (err, stdout, stderr) => {
      if (err) {
        console.error(err);
        return;
      }
      if (stderr) {
        console.error(stderr);
        return;
      }
      console.log(gras, `Fichier ${filename} créé`, resetColor);
    });
  }

  // MKDIR
  else if (input.startsWith('mkdir ')) {
    const directory = input.slice(6);
    exec(`mkdir ${directory}`, (err, stdout, stderr) => {
      if (err) {
        console.error(err);
        return;
      }
      if (stderr) {
        console.error(stderr);
      }
      console.log(stdout);
    });
  }

  // RM
  else if (input.startsWith('rm ')) {
    const path = input.slice(3);
    exec(`rm -r ${path}`, (err, stdout, stderr) => {
      if (err) {
        console.error(err);
        return;
      }
      if (stderr) {
        console.error(stderr);
        return;
      }
      console.log(stdout.trim());
    });
  }


  else 
  {
    command.stdin.write(input + '\n');
  }

  rl.setPrompt(prompt);
  rl.prompt();

});
