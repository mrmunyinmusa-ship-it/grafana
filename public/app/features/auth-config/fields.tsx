import { config } from '@grafana/runtime';
const clientSecretLabel = config.bootData.settings.authSecretLabel || 'Client secret';
