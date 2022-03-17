import clientsConfig from '../config/clients.config.json';

export function getLanguageFolder(language: string): string {
  return clientsConfig[language].folder;
}

export function getTestExtension(language: string): string | undefined {
  return clientsConfig[language]?.tests?.extension;
}

export function getTestOutputFolder(language: string): string | undefined {
  return clientsConfig[language]?.tests?.outputFolder;
}

export function getCustomGenerator(language: string): string | undefined {
  return clientsConfig[language]?.customGenerator;
}
