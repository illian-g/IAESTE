'use strict'; //strict mode

const {Contract} = require('fabric-contract-api');
// const {ChaincodeServer, shim} = require('fabric-shim');
// const fs = require('fs');
// const path = require('path');

// const chaincodeConfigPath = path.join(__dirname, 'chaincode.json');
// const chaincodeConfig = JSON.parse(fs.readFileSync(chaincodeConfigPath, 'utf8'));

class PushValuesContract extends Contract{
    async createAsset(ctx, assetId, voltage, current, time, date){
        
        const exists = await this.AssetExists(ctx, assetId); 
        //same class instance as createAsset, refers to contract which is an instance of PushValuesContract
        if (exists){
            throw new Error(`Asset ${assetId} already exists`);
        }

        const asset = {
            assetId: assetId,
            voltage: voltage,
            current: current,
            time: time,
            date: date
        };

        await ctx.stub.putState(assetId, Buffer.from(JSON.stringify(asset)));
        return `Asset ${assetId} created` ;
    }
    async AssetExists(ctx, assetId){
        const buffer = await ctx.stub.getState(assetId);
        return (!!buffer && buffer.length>0);
    }

    async ReadAsset(ctx, assetId) {
    const buffer = await ctx.stub.getState(assetId);
    if (!buffer || buffer.length === 0) {
        throw new Error(`Asset ${assetId} does not exist`);
    }
    return buffer.toString();
}

    
}
// const server= new ChaincodeServer({
//     chaincode: new PushValuesContract(),
//     chaincode_id: chaincodeConfig.chaincode_id,
//     tlsOptions: {
//         clientCert: Buffer.from(chaincodeConfig.client_cert),
//         clientKey: Buffer.from (chaincodeConfig.client_key),
//         rootCert: Buffer.from(chaincodeConfig.root_cert)
//     },
//     peerAddress: chaincodeConfig.peer_address,
//     identityContext: {
//         mspId: chaincodeConfig.mspid
//     },
//     peerEndpointOptions:{
//         host: chaincodeConfig.peer_address.split(':')[0],
//         port: parseInt(chaincodeConfig.peer_address.split(':')[1])
//     }

// });

module.exports = PushValuesContract;
// server.start();