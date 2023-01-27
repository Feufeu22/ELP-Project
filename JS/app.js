#!/usr/bin/env node

// **************** //
// ** JS PROJECT ** //
// **************** //

// ** MES FONCTIONS ** //

// Conversion base 10 vers 2
function de10a2(base10)
{
    liste = [];
    reste = base10;
    Lbase2 = [255,128,64,32,16,8,4,2,1];
    for(k=0;k<Lbase2.length;k++)
    {
        if(Lbase2[k]<=reste)
        {
            liste.push(1);
            reste = reste - Lbase2[k];
        }
        else
        {
            liste.push(0);
        } 
    }
    res=""
    for(i=0;i<liste.length;i++)
    {
        res += liste[i];
    }
    console.log(res);
}

// Conversion base 2 vers 10
const ExBase2 = [1,0,0,1,0,1];

function de2a10(base2)
{
    S = 0;
    for(i=0; i<base2.length ; i++)
        S+=base2[i]*2**(base2.length -i-1);
    return S;   
}


// ** MAIN ** //

console.log("\n#########################################");
console.log("#   CONVERTISSEUR BASE 10 / BASE 2   #");
console.log("#########################################\n");

const readline = require('readline');
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.prompt();

console.log("Bienvenue dans le convertisseur de base 10 vers base 2 et inversement !");
console.log("Si vous voulez quitter, tapez 'quit'");
console.log("\nQuelle type de conversion voulez-vous faire ?");
console.log("Taper '1' pour Base 10 vers Base 2");
console.log("Taper '2' pour Base 2 vers Base 10\n");

rl.on('line', (line) => {

  // Traitement de la commande
  if (line === 'quit') 
  {
    console.log('Au revoir !');
    rl.close();
    process.exit();
  } 

  else 
  { 
    if (line === '1')
    {
        console.log("Vous avez choisi de convertir de la base 10 vers la base 2");
        console.log("Votre nombre en base 10 :");
        rl.on('line', (line) => {
            if (line === 'quit') 
            {
                console.log('Au revoir !');
                rl.close();
                process.exit();
            } 
            else
            {
                console.log("Votre nombre en base 2 :");
                de10a2(line);
            }
        });
    }
    else if (line === '2')
    {
        console.log("Vous avez choisi de convertir de la base 2 vers la base 10");
        console.log("Votre nombre en base 2 :");
        rl.on('line', (line) => {
            if (line === 'quit') 
            {
                console.log('Au revoir !');
                rl.close();
                process.exit();
            } 
            else
            {
                console.log("Votre nombre en base 10 :");
                console.log(de2a10(line));
            }
        });
    }
  }

});

