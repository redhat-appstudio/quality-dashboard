import React, { FC, ReactElement, useState, useEffect } from 'react';
import { ReactReduxContext, useSelector } from 'react-redux';
import {
  Title,
  PageSectionVariants,
  PageSection,
  Icon,
  Text,
  Grid, GridItem,
  Toolbar, ToolbarItem, SearchInput
} from '@patternfly/react-core'
import { TableComposable, Thead, Tr, Th, Tbody, Td } from '@patternfly/react-table';
import CheckCircleIcon from '@patternfly/react-icons/dist/esm/icons/check-circle-icon';
import { listInstalledPlugins} from '@app/utils/APIService';

function importAll(r) {
  let images = {};
   r.keys().forEach((item, index) => { images[item.replace('./', '')] = r(item); });
  return images
}

type HubProps = {
  name: string,
  id: number,
  bio?: string,
}

interface PluginResponse {
  plugin: Plugin, 
  status:{
    installed: boolean
  }
}
type Plugin = {
  name: string,
  description: string,
  logo: string,
  category: string,
  reason: string,
}

const columnNames = {
  name: "Name",
  description: "Description",
  logo: "",
  category: "Category",
  status: "Status",
};

export const PInstalled: FC<HubProps> = ({/* destructured props */}): ReactElement => { 
  const images = importAll(require.context('../../images', false, /\.(png|jpe?g|svg)$/));
  const currentTeam = useSelector((state: any) => state.teams.Team);
  const widths = {
    default: '100px',
    sm: '80px',
    md: '150px',
    lg: '300px',
    xl: '350px',
    '2xl': '500px'
  };

  let [cards, setCards] = useState<Array<PluginResponse>>([])
  const [searchValue, setSearchValue] = React.useState('');
  const onChange = (value: string) => {
    setSearchValue(value);
  };

  const onFilter = (card: PluginResponse) => {
    if(!card.status.installed){
      return false
    }
    if (searchValue === '') {
      return true
    }

    let input: RegExp;
    try {
      input = new RegExp(searchValue, 'i');
    } catch (err) {
      input = new RegExp(searchValue.replace(/[.*+?^${}()|[\]\\]/g, '\\$&'), 'i');
    }
    return card.plugin.name.search(input) >= 0 

  };

  const filteredCards = cards.filter(onFilter);

  const listAllPlugins = () =>  {
    if(currentTeam == ''){
      console.error( "team is empty. cannot get plugins")
      return
    }
    listInstalledPlugins(currentTeam).then(res => {
      if(res.code == 200){
        setCards(res.data)
        console.log("response", res)
      } else {
        throw("Error getting plugins list")
      }
    })
  };

  useEffect(() => {
    if(currentTeam != ''){ listAllPlugins() }
  }, [currentTeam]);

  return (
    <React.Fragment>
      <PageSection variant={PageSectionVariants.light}>
        <Title headingLevel='h1'>Installed Plugins</Title>
        <Text style={{paddingTop: "1em"}}>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          Et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </Text>
      </PageSection>
      <Grid>
          <GridItem span={12} >
            <PageSection variant={PageSectionVariants.light}>
              <Toolbar id="toolbar-items-example">
                <ToolbarItem variant="search-filter" widths={widths}>
                  <SearchInput
                    placeholder="Find by name"
                    value={searchValue}
                    onChange={(_event, value) => onChange(value)}
                    onClear={() => onChange('')}
                  />
                </ToolbarItem>
              </Toolbar>
            </PageSection>
            <PageSection variant={PageSectionVariants.light}>
              <TableComposable aria-label="Sortable table custom toolbar">
                <Thead>
                  <Tr>
                    <Th></Th>
                    <Th>{columnNames.name}</Th>
                    <Th width={70}>{columnNames.description}</Th>
                    <Th>{columnNames.category}</Th>
                    <Th>{columnNames.status}</Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {filteredCards.filter(onFilter).map((product, rowIndex) => (
                    <Tr key={rowIndex}>
                      <Td style={{width:'50px'}}><img src={images[product.plugin.logo].default} width={50} style={{width:'50px'}}/></Td>
                      <Td style={{verticalAlign:'top'}} dataLabel={columnNames.name}>{product.plugin.name}</Td>
                      <Td style={{verticalAlign:'top'}} dataLabel={columnNames.description}>{product.plugin.description}</Td>
                      <Td style={{verticalAlign:'top', textTransform: 'capitalize'}} dataLabel={columnNames.category}>{product.plugin.category}</Td>
                      <Td style={{verticalAlign:'top', textTransform: 'capitalize'}} dataLabel={columnNames.status}>
                        { product.status.installed && <span><Icon status="success"><CheckCircleIcon /></Icon> &nbsp; Installed </span> }
                      </Td>
                    </Tr>
                  ))}
                </Tbody>
              </TableComposable>
            </PageSection>
        </GridItem>
      </Grid>
    </React.Fragment>
  )
};