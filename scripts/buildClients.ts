import { CI, run } from './common';
import { getClientsConfigField, getLanguageFolder } from './config';
import { createSpinner } from './oraLog';
import type { Generator } from './types';

const multiBuildLanguage = new Set(['javascript']);

/**
 * Build only a specific client for one language, used by javascript for example.
 */
async function buildPerClient(
  { language, key, additionalProperties }: Generator,
  verbose: boolean
): Promise<void> {
  const spinner = createSpinner(`building ${key}`, verbose).start();
  const npmNamespace = getClientsConfigField(language, 'npmNamespace');

  switch (language) {
    case 'javascript':
      await run(
        `yarn workspace ${npmNamespace}/${additionalProperties.packageName} clean`,
        {
          verbose,
          cwd: getLanguageFolder(language),
        }
      );
      await run(
        `SKIP_UTILS=true yarn build ${npmNamespace}/${additionalProperties.packageName}`,
        { verbose, cwd: getLanguageFolder(language) }
      );
      break;
    default:
  }
  spinner.succeed();
}

/**
 * Build all client for a language at the same time, for those who live in the same folder.
 */
async function buildAllClients(
  language: string,
  verbose: boolean
): Promise<void> {
  const spinner = createSpinner(`building '${language}'`, verbose).start();
  switch (language) {
    case 'java':
      await run(
        `./gradle/gradlew --no-daemon -p ${getLanguageFolder(
          language
        )} assemble`,
        {
          verbose,
        }
      );
      break;
    case 'php':
      break;
    default:
  }
  spinner.succeed();
}

export async function buildClients(
  generators: Generator[],
  { verbose, skipUtils }: { verbose: boolean; skipUtils: boolean }
): Promise<void> {
  const langs = [...new Set(generators.map((gen) => gen.language))];

  if (langs.includes('javascript')) {
    await run('yarn install', {
      verbose,
      cwd: getLanguageFolder('javascript'),
    });
  }

  if (!CI && !skipUtils && langs.includes('javascript')) {
    const spinner = createSpinner(
      "building 'JavaScript' utils",
      verbose
    ).start();

    await run('yarn clean:utils', {
      verbose,
      cwd: getLanguageFolder('javascript'),
    });
    await run('yarn build:utils', {
      verbose,
      cwd: getLanguageFolder('javascript'),
    });

    spinner.succeed();
  }

  await Promise.all([
    ...generators
      .filter(({ language }) => multiBuildLanguage.has(language))
      .map((gen) => buildPerClient(gen, verbose)),
    ...langs
      .filter((lang) => !multiBuildLanguage.has(lang))
      .map((lang) => buildAllClients(lang, verbose)),
  ]);
}
